package templating

import (
	"embed"
	"html/template"
	"io"
)

//go:embed "templates/*"
var templates embed.FS

func RenderPost(w io.Writer, p Post) error {
	templ, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "post.gohtml", p); err != nil {
		return err
	}

	return nil
}
