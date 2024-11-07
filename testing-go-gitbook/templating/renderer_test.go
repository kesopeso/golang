package templating_test

import (
	"bytes"
	"templating"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderPost(t *testing.T) {
	post := templating.Post{
		Title:       "Test post",
		Description: "Testing rendering",
		Tags:        []string{"rendering", "tests"},
		Body: `# Body title

This is some body to test the render with

## Subtitle

Another body paragraph.`,
	}

	postRenderer, err := templating.NewPostRenderer()
	if err != nil {
		t.Fatal("error should not occur", err)
	}
	buf := bytes.Buffer{}
	err = postRenderer.RenderPost(&buf, post)
	if err != nil {
		t.Fatal("error should not occur", err)
	}

	approvals.VerifyString(t, buf.String())
}

func TestRenderIndex(t *testing.T) {
	posts := []templating.Post{
		{Title: "Post 1"},
		{Title: "Post 2"},
		{Title: "Post 3"},
	}

	postRenderer, err := templating.NewPostRenderer()
	if err != nil {
		t.Fatal("error should not occur", err)
	}
	buf := bytes.Buffer{}
	err = postRenderer.RenderIndex(&buf, posts)
	if err != nil {
		t.Fatal("error should not occur", err)
	}

	approvals.VerifyString(t, buf.String())
}
