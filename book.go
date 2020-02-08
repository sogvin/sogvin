package sogvin

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gregoryv/sogvin/internal"
	. "github.com/gregoryv/web"
)

func NewBook() *Book {
	book := Book{}
	toc(&book)
	return &book
}

type Book struct {
	pages []*Page
}

// Saves all pages and table of contents
func (book *Book) SaveTo(base string) error {
	for _, page := range book.pages {
		page.SaveTo(base)
	}
	return nil
}

func findH1(article *Element) string {
	var buf bytes.Buffer
	w := NewHtmlWriter(&buf)
	w.WriteHtml(article)
	from := bytes.Index(buf.Bytes(), []byte("<h1>")) + 4
	to := bytes.Index(buf.Bytes(), []byte("</h1>"))
	return strings.TrimSpace(string(buf.Bytes()[from:to]))
}

func (book *Book) AddPage(right string, article *Element) *Element {
	filename := filenameFrom(findH1(article)) + ".html"
	page := newPage(
		filename,
		header(right+" - "+A(Href("index.html"), "Software Engineering").String()),
		article,
		footer,
	)
	book.pages = append(book.pages, page)
	return linkToPage(page)
}

func linkToPage(page *Page) *Element {
	return Li(A(Href(page.Filename), findH1(page.Element)))
}

func newPage(filename string, header, article, footer *Element) *Page {
	return NewPage(filename,
		Html(en,
			Head(utf8, viewport, theme, a4),
			Body(header, article, footer),
		),
	)
}

func filenameFrom(in string) string {
	tidy := bytes.NewBufferString("")
	var inside bool
	for _, c := range in {
		switch c {
		case '(', ')':
			continue
		case ' ':
			tidy.WriteRune('_')
		case '<':
			inside = true
		case '>':
			inside = false
		default:
			if inside {
				continue
			}
			tidy.WriteString(strings.ToLower(string(c)))
		}
	}
	return tidy.String()
}

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

func header(right string) *Element {
	h := Header()
	if right != "" {
		h = h.With(Code(right))
	}
	return h
}

func stylesheet(href string) *Element {
	return Link(Rel("stylesheet"), Type("text/css"), Href(href))
}

func boxnote(txt string, cm float64) *Element {
	return Div(Class("boxnote"),
		&Attribute{
			Name: "style",
			Val:  fmt.Sprintf("margin-top: %vcm", cm),
		},
		txt,
	)
}

func loadGoFile(filename string, from, to int) *Element {
	return goCode(internal.LoadFile(filename, from, to))
}

func goCode(v string) *Element {
	return Pre(Class("srcfile"), Code(Class("go"), v))
}

func gregoryv(name, txt string) *Element {
	return Li(
		fmt.Sprintf(
			`<a href="https://github.com/gregoryv/%s">%s</a> - %s`,
			name, name, txt,
		),
	)
}

func shellCommand(v string) *Element {
	return Pre(Class("command"), Code(v))
}
