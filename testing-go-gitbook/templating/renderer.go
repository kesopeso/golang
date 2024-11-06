package templating

import (
	"fmt"
	"io"
)

func RenderPost(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
