package notes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	. "github.com/gregoryv/web/doctype"
)

func NewBook() *Book {
	return &Book{
		pages: []*Page{
			NewPage(PurposeOfFuncMain, "func main()", "purpose_of_func_main.html"),
			NewPage(NexusPattern, "Nexus pattern", "nexus_pattern.html"),
			NewPage(InlineTestHelpers, "Testing", "inline_test_helpers.html"),
			NewPage(GracefulServerShutdown, "Shutdown", "graceful_server_shutdown.html"),

			//NewPage(Dictionary, "Dictionary", "dictionary.html"),
		},
	}
}

type Book struct {
	pages []*Page
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
	index := NewPage(art, "", "index.html")
	index.SaveTo(base)
}

func findH1(article writerTo) string {
	var buf bytes.Buffer
	article.WriteTo(&buf)
	from := bytes.Index(buf.Bytes(), []byte("<h1>")) + 4
	to := bytes.Index(buf.Bytes(), []byte("</h1>"))
	return strings.TrimSpace(string(buf.Bytes()[from:to]))
}

func (page *Page) SaveTo(base string) {
	out := path.Join(base, page.filename)
	fmt.Println("  ", out)
	fh, _ := os.Create(out)
	page.html.WriteTo(fh)
	fh.Close()
}

type writerTo interface {
	WriteTo(io.Writer) (int, error)
}

func NewPage(article *Tag, right, filename string) *Page {
	return &Page{
		html: Html(en,
			Head(utf8, viewport, theme, a4),
			Body(
				header("", right),
				article,
				footer,
			),
		),
		right:    right,
		filename: filename,
	}
}

type Page struct {
	html     *HtmlTag
	right    string
	filename string
}
