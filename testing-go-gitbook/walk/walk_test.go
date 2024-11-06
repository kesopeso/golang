package walk_test

import (
	"reflect"
	"testing"
	"walk"
)

func TestWalkAlt(t *testing.T) {
	t.Run("simple cases", func(t *testing.T) {
		testCases := []struct {
			name  string
			input any
			want  []string
		}{
			{"simple string",
				"just testing",
				[]string{"just testing"},
			},
			{
				"single string struct",
				struct{ Value string }{"some string"},
				[]string{"some string"},
			},
			{
				"pointer struct",
				struct {
					Value *struct {
						InnerValue   string
						AnotherValue string
					}
				}{
					&struct {
						InnerValue   string
						AnotherValue string
					}{"inner value", "another value"},
				},
				[]string{"inner value", "another value"},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := make([]string, 0, 1)
				walk.WalkAlt(tc.input, func(v string) {
					got = append(got, v)
				})
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("slices mismatch, got %v, want %v", got, tc.want)
				}
			})
		}
	})

	t.Run("channels test", func(t *testing.T) {
		ch := make(chan string)

		want := []string{"value1", "value2"}
		got := make([]string, 0, 2)

		go func() {
			for _, v := range want {
				ch <- v
			}
			close(ch)
		}()

		walk.WalkAlt(ch, func(v string) {
			got = append(got, v)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("chan result mismatch, want %v, got %v", want, got)
		}
	})
}

func TestWalk(t *testing.T) {
	t.Run("returns slice of string field names", func(t *testing.T) {
		obj := struct {
			Name     string
			Surname  string
			Age      int
			Nickname string
			Profile  struct {
				Address string
				Count   int
			}
			Alt *struct {
				Field string
			}
		}{"name", "surname", 5, "nickname", struct {
			Address string
			Count   int
		}{"address", 5}, &struct {
			Field string
		}{"alt"}}

		type result struct {
			name  string
			value string
		}

		want := []result{
			{"Name", "name"},
			{"Surname", "surname"},
			{"Nickname", "nickname"},
			{"Profile.Address", "address"},
			{"Alt.Field", "alt"},
		}

		got := make([]result, 0, 3)
		walk.Walk(obj, func(name string, value string) {
			got = append(got, result{name, value})
		}, "")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("returns empty slice for non struct variable", func(t *testing.T) {
		walk.Walk(5, func(name string, value string) {
			t.Errorf("function should not be called, received name %s, value %s", name, value)
		}, "")
	})
}
