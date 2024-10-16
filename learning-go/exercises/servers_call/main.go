package main

import (
	"context"
	"fmt"
	"time"
)

type AIn string
type AOut string
type BIn string
type BOut string
type CIn string
type COut string

func processA(data AIn) (AOut, error) {
	time.Sleep(10 * time.Millisecond)
	return AOut(string(data) + string(data)), nil
}

func processB(data BIn) (BOut, error) {
	time.Sleep(25 * time.Millisecond)
	return BOut(string(data) + string(data) + string(data)), nil
}

func processC(data CIn) (COut, error) {
	time.Sleep(23 * time.Millisecond)
	return COut(string(data) + string(data) + string(data) + string(data)), nil
}

type abProcessor struct {
	chA  chan AOut
	chB  chan BOut
	errs chan error
}

type cProcessor struct {
	chC chan COut
	err chan error
}

type Processor struct {
	abProcessor *abProcessor
	cProcessor  *cProcessor
}

func getAbProcessor() *abProcessor {
	return &abProcessor{
		chA:  make(chan AOut, 1),
		chB:  make(chan BOut, 1),
		errs: make(chan error, 2),
	}
}

func (abp *abProcessor) Process(dataA AIn, dataB BIn) {
	go func() {
		result, err := processA(dataA)
		if err != nil {
			abp.errs <- err
			return
		}
		abp.chA <- result
	}()

	go func() {
		result, err := processB(dataB)
		if err != nil {
			abp.errs <- err
			return
		}
		abp.chB <- result
	}()
}

func (abp *abProcessor) Wait(ctx context.Context) (AOut, BOut, error) {
	var resultA AOut
	var resultB BOut
	var resultErr error

	for count := 0; count < 2; count++ {
		select {
		case resultA = <-abp.chA:
		case resultB = <-abp.chB:
		case resultErr = <-abp.errs:
			var aZero AOut
			var bZero BOut
			return aZero, bZero, resultErr
		case <-ctx.Done():
			var aZero AOut
			var bZero BOut
			return aZero, bZero, ctx.Err()
		}
	}

	return resultA, resultB, nil
}

func getCProcessor() *cProcessor {
	return &cProcessor{
		chC: make(chan COut, 1),
		err: make(chan error, 1),
	}
}

func (cP *cProcessor) Process(dataC CIn) {
	go func() {
		result, err := processC(dataC)
		if err != nil {
			cP.err <- err
			return
		}
		cP.chC <- result
	}()
}

func (cP *cProcessor) Wait(ctx context.Context) (COut, error) {
	select {
	case result := <-cP.chC:
		return result, nil
	case err := <-cP.err:
		var zero COut
		return zero, err
	case <-ctx.Done():
		var zero COut
		return zero, ctx.Err()
	}
}

func GetProcessor() *Processor {
	return &Processor{
		abProcessor: getAbProcessor(),
		cProcessor:  getCProcessor(),
	}
}

func (p *Processor) Process(dataA AIn, dataB BIn, timeout time.Duration) (COut, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	p.abProcessor.Process(dataA, dataB)
	aOut, bOut, err := p.abProcessor.Wait(ctx)
	if err != nil {
		var zero COut
		return zero, err
	}

	cIn := CIn(string(aOut) + string(bOut))
	p.cProcessor.Process(cIn)
	cOut, err := p.cProcessor.Wait(ctx)
	if err != nil {
		var zero COut
		return zero, err
	}

	return cOut, nil
}

func main() {
	processor := GetProcessor()
	result, err := processor.Process(AIn("a"), BIn("b"), 45*time.Millisecond)
	if err != nil {
		fmt.Println("Error occured while processing:", err)
		return
	}
	fmt.Println("Processing was successful. This is the result:", result)
}
