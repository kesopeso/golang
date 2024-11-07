package templating

import (
	"embed"
	"html/template"
	"io"
)

//go:embed "templates/*"
var templates embed.FS

type PostRenderer struct {
	t *template.Template
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.t.ExecuteTemplate(w, "post.gohtml", p)
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ}, nil
}
