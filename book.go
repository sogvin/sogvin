package notes

import (
	"bytes"
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
func (book *Book) SaveTo(base string) error {
	toc := Ul(Class("toc"))
	for _, p := range book.pages {
		p.SaveTo(base)
		toc = toc.With(
			Li(
				A(
					Href(p.filename),
					findH1(p.Element),
				),
			),
		)
	}
	art := Article(
		H1("Software Engineering"),
		P("Notes by", myname),
		H2("Table of Contents"),
		toc,
		H3("Design"),
		Ul(Class("toc"),
			gregoryv("draw", "software engineering diagrams"),
			gregoryv("web", "html generation"),
		),

		H3("Test"),
		Ul(Class("toc"),
			gregoryv("golden", "simplify use of golden files"),
			gregoryv("qual", "quality constraints"),
			gregoryv("ex", "indented JSON or redirect handler response to stdout"),
			gregoryv("uncover", "generate coverage reports from cover profiles"),
		),

		H3("Build"),
		Ul(Class("toc"),
			gregoryv("stamp", "parse build information to embed into your binary"),
			gregoryv("find", "search for files by name or content"),
		),
	)
	index := newPage(art, "", "index.html")
	return index.SaveTo(base)
}

func findH1(article *Element) string {
	var buf bytes.Buffer
	w := NewHtmlWriter(&buf)
	w.WriteHtml(article)
	from := bytes.Index(buf.Bytes(), []byte("<h1>")) + 4
	to := bytes.Index(buf.Bytes(), []byte("</h1>"))
	return strings.TrimSpace(string(buf.Bytes()[from:to]))
}

func NewPageA4(article *Element, right, filename string) *PageA4 {
	return newPage(article, right+" - Software Engineering", filename)
}

func newPage(article *Element, right, filename string) *PageA4 {
	return &PageA4{
		Page: NewPage(filename, Html(en,
			Head(utf8, viewport, theme, a4),
			Body(
				header(right),
				article,
				footer,
			),
		)),
		right:    right,
		filename: filename,
	}
}

type PageA4 struct {
	*Page
	right    string
	filename string
}
