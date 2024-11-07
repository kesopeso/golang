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

type postViewModel struct {
	Post
	BodyHTML template.HTML
}

type postsIndexViewModel struct {
	Title      string
	LinkSuffix string
}

type PostRenderer struct {
	t *template.Template
	p *parser.Parser
	r *html.Renderer
}

func (r *PostRenderer) RenderPost(w io.Writer, p Post) error {
	return r.t.ExecuteTemplate(w, "post.gohtml", r.newPostViewModel(p))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.t.ExecuteTemplate(w, "index.gohtml", r.newPostsIndexViewModel(posts))
}

func (r *PostRenderer) newPostViewModel(p Post) postViewModel {
	return postViewModel{Post: p, BodyHTML: convertMarkdownToHTML(r, p)}
}

func (r *PostRenderer) newPostsIndexViewModel(posts []Post) []postsIndexViewModel {
	result := make([]postsIndexViewModel, len(posts))
	for i, p := range posts {
		result[i] = postsIndexViewModel{p.Title, createLinkFromTitle(p.Title)}
	}
	return result
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return &PostRenderer{templ, p, renderer}, nil
}

func convertMarkdownToHTML(r *PostRenderer, p Post) template.HTML {
	bHTML := markdown.ToHTML([]byte(p.Body), r.p, r.r)
	html := template.HTML(strings.TrimSuffix(string(bHTML), "\n"))
	bodyHTML := template.HTML(html)
	return bodyHTML
}

func createLinkFromTitle(title string) string {
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")
	return title
}
