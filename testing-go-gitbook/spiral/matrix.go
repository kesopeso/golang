package spiral

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type SpiralMatrix struct {
	matrixSize      int
	totalIterations int
	totalCircles    int
	matrix          [][]bool
}

func (sm *SpiralMatrix) Write(p []byte) (n int, err error) {
	sm.matrix = make([][]bool, sm.matrixSize)
	for i := range sm.matrix {
		sm.matrix[i] = make([]bool, sm.matrixSize)
	}

Outer:
	for _, pointData := range strings.Split(string(p), "\n") {
		pointXY := strings.Split(pointData, "#")
		if len(pointXY) != 2 {
			err = errors.New("corrupt input data")
			break Outer
		}

		pointX, err := strconv.ParseFloat(pointXY[0], 64)
		if err != nil {
			break Outer
		}
		pointX += (float64(sm.matrixSize) / 2)

		pointY, err := strconv.ParseFloat(pointXY[1], 64)
		if err != nil {
			break Outer
		}
		pointY = -pointY + (float64(sm.matrixSize) / 2)

		xIdx := int(math.Floor(pointX)) // column
		yIdx := int(math.Floor(pointY)) // row
		if yIdx > sm.matrixSize-1 || xIdx > sm.matrixSize-1 {
			err = errors.New(fmt.Sprintf("matrix entry out of bounds x = %d, y = %d, matrix size %d", xIdx, yIdx, sm.matrixSize))
			break Outer
		}

		sm.matrix[yIdx][xIdx] = true
	}

	if err == nil {
		n = len(p)
	}
	return
}

func (sm *SpiralMatrix) Print(w io.Writer) error {
	startingR := float64(sm.matrixSize) / 2
	_, err := WriteSpiral(sm, startingR, sm.totalIterations, sm.totalCircles)
	if err != nil {
		return err
	}

	for _, row := range sm.matrix {
		for _, column := range row {
			if column {
				fmt.Fprint(w, "#")
			} else {
				fmt.Fprint(w, " ")
			}
		}
		fmt.Fprint(w, "\n")
	}

	return nil
}

func NewSpiralMatrix(matrixSize, totalIterations, totalCircles int) *SpiralMatrix {
	return &SpiralMatrix{matrixSize, totalIterations, totalCircles, nil}
}
