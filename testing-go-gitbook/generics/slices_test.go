package generics_test

import (
	"generics"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	data := []int{1, 2, 3, 4}
	got := generics.Sum(data)
	want := 10
	if got != want {
		t.Errorf("result mismatch, got %d, want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	slice1 := []int{5, 5, 5}
	slice2 := []int{10, 10, 10}
	got := generics.SumAllTails(slice1, slice2)
	want := []int{10, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("result mismatch, got %v, want %v", got, want)
	}
}
func TestMultiply(t *testing.T) {
	data := []int{1, 2, 3, 4}
	got := generics.Multiply(data)
	want := 24
	if got != want {
		t.Errorf("result mismatch, got %d, want %d", got, want)
	}
}

func TestReduce(t *testing.T) {
	data := []string{"a", "b", "c"}
	got := generics.Reduce(data, func(result string, current string) string {
		return result + current
	}, "")
	want := "abc"
	if got != want {
		t.Errorf("got %v, want %v mismatch", got, want)
	}
}
