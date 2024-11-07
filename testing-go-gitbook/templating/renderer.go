package templating

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed "templates/*"
var templates embed.FS

type PostHTML struct {
	Post
	Body template.HTML
}

type PostRenderer struct {
	t *template.Template
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.t.ExecuteTemplate(w, "post.gohtml", newPostHTML(p))
}

func newPostHTML(p Post) PostHTML {
	return PostHTML{
		Post: p,
		Body: convertMarkdownToHTML(p.Body),
	}
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ}, nil
}

func convertMarkdownToHTML(md string) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	bHTML := markdown.ToHTML([]byte(md), p, renderer)
	html := template.HTML(strings.TrimSuffix(string(bHTML), "\n"))

	return html
}
