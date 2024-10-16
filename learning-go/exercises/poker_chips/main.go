package main

import (
	"fmt"
	"time"
)

const elementsCount = 5
const startPrime = 400
const primesCount = 20
const maxNumberOfChips = 5

func main() {
	start := time.Now()

	primes := getPrimes()
	fmt.Println("primes", primes)

	combinations := getCombinationsRecursive([][elementsCount]int{}, primes, []int{}, 0, elementsCount)
	noncollidingCombination, ok := getFirstNoncollidingCombination(combinations, maxNumberOfChips)
	if !ok {
		fmt.Println("no such non colliding combination")
	} else {
		fmt.Println("found non colliding combination", noncollidingCombination)
	}

	duration := time.Since(start)
	fmt.Println("duration", duration)
}

func getFirstNoncollidingCombination(combinations [][elementsCount]int, maxNum int) ([elementsCount]int, bool) {
	for _, combination := range combinations {
		resultHashSet := make(map[int][elementsCount]int)
		isNoncollidingCombination := true

		fmt.Println("checking combination", combination)

	startLoop:
		for i := 0; i <= maxNum; i++ {
			for j := 0; j <= maxNum; j++ {
				for k := 0; k <= maxNum; k++ {
					for l := 0; l <= maxNum; l++ {
						for m := 0; m <= maxNum; m++ {
							result := i*combination[0] + j*combination[1] + k*combination[2] + l*combination[3] + m*combination[4]
							if result == 0 {
								continue
							}
							colliding, found := resultHashSet[result]
							if found {
								fmt.Println("found colliding combination", combination, colliding, [elementsCount]int{i, j, k, l, m})
								isNoncollidingCombination = false
								break startLoop
							}
							resultHashSet[result] = [elementsCount]int{i, j, k, l, m}
						}
					}
				}
			}
		}

		if isNoncollidingCombination {
			return combination, true
		}
	}

	var zero [elementsCount]int
	return zero, false
}

func getCombinationsRecursive(partialResult [][elementsCount]int, numbers []int, currentCombination []int, startIdx int, endOffset int) [][elementsCount]int {
	endOffset = endOffset - 1
	returnResult := endOffset == 0 && len(currentCombination) == elementsCount-1

	for i := startIdx; i < len(numbers)-endOffset; i++ {
		currentCombination := append(currentCombination, numbers[i])
		if returnResult {
			partialResult = append(partialResult, [elementsCount]int(currentCombination))
			continue
		}
		partialResult = getCombinationsRecursive(partialResult, numbers, currentCombination, i+1, endOffset)
	}

	return partialResult
}

func getPrimes() []int {
	if primesCount < elementsCount {
		panic("count must be at least " + string(elementsCount))
	}

	found := 0
	primes := make([]int, 0, primesCount)

	i := startPrime

	for {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
			found++

			if found == primesCount {
				return primes
			}
		}

		i++
	}
}
