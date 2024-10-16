package file_parser_test

import (
	"bytes"
	"ch15/file_parser"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FuzzParseData(f *testing.F) {
	f.Add([]byte("3\nhello\ngoodbye\ngreetings"))
	f.Add([]byte("0\n"))

	f.Fuzz(func(t *testing.T, in []byte) {
		r := bytes.NewReader(in)
		out, err := file_parser.ParseData(r, 1000)
		if err != nil {
			t.Skip("handled error")
		}
		roundTrip := file_parser.ToData(out)
		rtr := bytes.NewReader(roundTrip)
		out2, err := file_parser.ParseData(rtr, 1000)
		if diff := cmp.Diff(out, out2); diff != "" {
			t.Error(diff)
		}
	})
}

func TestParseData(t *testing.T) {
	data := []struct {
		name   string
		in     []byte
		out    []string
		errMsg string
	}{
		{
			"simple",
			[]byte("3\nhello\ngoodbye\ngreetings"),
			[]string{"hello", "goodbye", "greetings"},
			"",
		},
		{
			"too_many_lines",
			[]byte("4\nhello\ngoodbye\ngreetings\naloha"),
			nil,
			"too many lines",
		},
		{
			"no_empty_lines",
			[]byte("1\n\r\r"),
			nil,
			"no empty lines",
		},
		{
			"empty_error",
			[]byte(""),
			nil,
			"empty",
		},
		{
			"zero",
			[]byte("0\n"),
			[]string{},
			"",
		},
		{
			"bad_number",
			[]byte("no number\n"),
			nil,
			`strconv.Atoi: parsing "no number": invalid syntax`,
		},
		{
			"too_few_lines",
			[]byte("2\none liner"),
			nil,
			"too few lines",
		},
		{
			"negative_count",
			[]byte("-8\n"),
			nil,
			"no negative numbers",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			r := bytes.NewReader(d.in)
			out, err := file_parser.ParseData(r, 3)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if diff := cmp.Diff(d.out, out); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(d.errMsg, errMsg); diff != "" {
				t.Error(diff)
			}
		})
	}
}
