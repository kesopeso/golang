// tools for code coverage
// go test -v -cover -coverprofile=c.out
// go tool cover -html=c.out
package cmp_test

import (
	"ch15/cmp"
	"testing"

	goCmp "github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	want := cmp.Person{
		Name: "Dennis",
		Age:  30,
	}

	got := cmp.CreatePerson("Dennis", 30)

	comparer := goCmp.Comparer(func(p1 cmp.Person, p2 cmp.Person) bool {
		return p1.Name == p2.Name && p1.Age == p2.Age
	})

	if diff := goCmp.Diff(want, got, comparer); diff != "" {
		t.Error("person structs dont match", diff)
	}
}

func TestAdd(t *testing.T) {
	data := []struct {
		i    int
		j    int
		want int
	}{
		{1, 1, 2},
		{5, 10, 15},
		{100, 10, 110},
	}

	for _, d := range data {
		t.Run("Add", func(t *testing.T) {
			t.Parallel()
			got := cmp.Add(d.i, d.j)
			if got != d.want {
				t.Errorf("addition not working, i: %d, j: %d, got: %d, want: %d", d.i, d.j, got, d.want)
			}
		})
	}
}
