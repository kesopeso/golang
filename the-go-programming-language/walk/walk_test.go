package walk_test

import (
	"reflect"
	"testing"
	"walk"
)

func TestWalk(t *testing.T) {
	t.Run("returns slice of string field names", func(t *testing.T) {
		obj := struct {
			Name     string
			Surname  string
			Age      int
			Nickname string
		}{"name", "surname", 5, "nickname"}

		type result struct {
			name  string
			value string
		}

		want := []result{
			{"Name", "name"},
			{"Surname", "surname"},
			{"Nickname", "nickname"},
		}

		got := make([]result, 0, 3)
		walk.Walk(obj, func(name string, value string) {
			got = append(got, result{name, value})
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("returns empty slice for non struct variable", func(t *testing.T) {
		walk.Walk(5, func(name string, value string) {
			t.Errorf("function should not be called, received name %s, value %s", name, value)
		})
	})
}
