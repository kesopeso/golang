package generics

func Find[T any](data []T, predicate func(element T) bool) (T, bool) {
	for _, d := range data {
		if predicate(d) {
			return d, true
		}
	}
	var zero T
	return zero, false
}
