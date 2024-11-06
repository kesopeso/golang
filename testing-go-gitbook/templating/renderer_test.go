package templating_test

import (
	"bytes"
	"templating"
	"testing"
)

func TestRender(t *testing.T) {
	post := templating.Post{
		Title:       "Test post",
		Description: "Testing rendering",
		Tags:        []string{"rendering", "tests"},
		Body:        "This is some body to test the render with",
	}

	t.Run("render post", func(t *testing.T) {
		buf := bytes.Buffer{}

		err := templating.RenderPost(&buf, post)
		if err != nil {
			t.Fatal("error should not occur", err)
		}

		got := buf.String()
		want := "<h1>Test post</h1>"

		if got != want {
			t.Errorf("title missmatch, got %s, want %s", got, want)
		}
	})
}
