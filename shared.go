package notes

import (
	"fmt"

	"github.com/gregoryv/notes/internal"
	. "github.com/gregoryv/web"
)

var (
	en       = Lang("en")
	utf8     = Meta(Charset("utf-8"))
	viewport = Meta(
		Name("viewport"),
		Content("width=device-width, initial-scale=1.0"),
	)
	theme  = stylesheet("theme.css")
	a4     = stylesheet("a4.css")
	footer = Footer(myname)
	myname = "Gregory Vin&ccaron;i&cacute;"
)

func header(right string) *Tag {
	h := Header()
	if right != "" {
		h = h.With(Code(right))
	}
	return h
}

func stylesheet(href string) *Tag {
	return Link(
		Rel("stylesheet"),
		Type("text/css"),
		Href(href),
	)
}

func boxnote(txt string, cm float64) *Tag {
	return Div(Class("boxnote"),
		&Attr{Name: "style", Val: fmt.Sprintf("margin-top: %vcm", cm)},
		txt,
	)
}

func loadGoFile(filename string, from, to int) string {
	return `<pre class="srcfile"><code class="go">` +
		internal.LoadFile(filename, from, to) + "</code></pre>"
}