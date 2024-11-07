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
		Body: `# Body title

This is some body to test the render with

## Subtitle

Another body paragraph.`,
	}

	t.Run("render post", func(t *testing.T) {
		buf := bytes.Buffer{}

		postRenderer, err := templating.NewPostRenderer()
		if err != nil {
			t.Fatal("error should not occur", err)
		}
		err = postRenderer.Render(&buf, post)
		if err != nil {
			t.Fatal("error should not occur", err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
