package generics

func Sum(numbers []int) int {
	return Reduce(numbers, func(sum int, current int) int {
		return sum + current
	}, 0)

	// var sum int
	// for _, v := range numbers {
	// 	sum += v
	// }
	// return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	return Reduce(numbersToSum, func(sum []int, current []int) []int {
		if len(current) == 0 {
			return append(sum, 0)
		}
		return append(sum, Sum(current[1:]))
	}, make([]int, 0))

	// var sums []int
	// for _, v := range numbersToSum {
	// 	if len(v) == 0 {
	// 		sums = append(sums, 0)
	// 		continue
	// 	}
	// 	sums = append(sums, Sum(v[1:]))
	// }
	// return sums
}

func Multiply(numbers []int) int {
	f := func(result int, current int) int {
		return result * current
	}
	return Reduce(numbers, f, 1)
}
