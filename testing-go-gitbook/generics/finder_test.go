package generics_test

import (
	"generics"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("element exists", func(t *testing.T) {
		got, found := generics.Find([]int{1, 2, 3}, func(x int) bool {
			return x == 2
		})
		want := 2

		if !found {
			t.Error("value should be found, but was not")
		}

		if got != want {
			t.Errorf("wrong number found, got %d, want %d", got, want)
		}
	})

	t.Run("element not found, return default value", func(t *testing.T) {
		got, found := generics.Find([]int{1, 2, 3}, func(x int) bool {
			return x == 4
		})

		if found {
			t.Error("value shouldn't be found, but it was")
		}

		want := 0
		if got != want {
			t.Errorf("wrong default returned, got %d, want %d", got, want)
		}
	})
}
