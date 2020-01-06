package notes

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

	. "github.com/gregoryv/web"
)

func NewBook() *Book {
	return &Book{
		pages: []*PageA4{
			NewPageA4(PurposeOfFuncMain, "func main()", "purpose_of_func_main.html"),
			NewPageA4(NexusPattern, "Nexus pattern", "nexus_pattern.html"),
			NewPageA4(InlineTestHelpers, "Testing", "inline_test_helpers.html"),
			NewPageA4(GracefulServerShutdown, "Shutdown", "graceful_server_shutdown.html"),

			//NewPage(Dictionary, "Dictionary", "dictionary.html"),
		},
	}
}

type Book struct {
	pages []*PageA4
}

// Saves all pages and table of contents
func (book *Book) SaveTo(base string) {
	toc := Ul(Class("toc"))
	for _, p := range book.pages {
		p.SaveTo(base)
		toc = toc.With(
			Li(
				A(
					Href(p.filename),
					findH1(p.html),
				),
			),
		)
	}
	art := Article(
		H1("Software Engineering"),
		P("Notes by", myname),
		H2("Table of Contents"),
		toc,
	)
	index := newPage(art, "", "index.html")
	index.SaveTo(base)
}

func findH1(article *Element) string {
	var buf bytes.Buffer
	w := NewHtmlWriter(&buf)
	w.WriteHtml(article)
	from := bytes.Index(buf.Bytes(), []byte("<h1>")) + 4
	to := bytes.Index(buf.Bytes(), []byte("</h1>"))
	return strings.TrimSpace(string(buf.Bytes()[from:to]))
}

func (page *PageA4) SaveTo(base string) {
	out := path.Join(base, page.filename)
	fmt.Println("  ", out)
	fh, _ := os.Create(out)
	w := NewHtmlWriter(fh)
	w.WriteHtml(page.html)
	fh.Close()
}

func NewPageA4(article *Element, right, filename string) *PageA4 {
	return newPage(article, right+" - Software Engineering", filename)
}

func newPage(article *Element, right, filename string) *PageA4 {
	return &PageA4{
		html: Html(en,
			Head(utf8, viewport, theme, a4),
			Body(
				header(right),
				article,
				footer,
			),
		),
		right:    right,
		filename: filename,
	}
}

type PageA4 struct {
	html     *Element
	right    string
	filename string
}
