package generics

func Reduce[T1, T2 any](data []T1, reduce func(T2, T1) T2, initialValue T2) T2 {
	result := initialValue
	for _, v := range data {
		result = reduce(result, v)
	}
	return result
}
