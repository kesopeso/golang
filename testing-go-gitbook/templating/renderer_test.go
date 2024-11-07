package templating_test

import (
	"bytes"
	"templating"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
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

		approvals.VerifyString(t, buf.String())
	})
}
