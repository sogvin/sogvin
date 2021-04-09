package sogvin

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

type Book struct {
	Title  string
	Author string
	pages  []*Page
	themes []*CSS
}

// Saves all pages and table of contents
func (book *Book) SaveTo(base string) error {
	for _, page := range book.pages {
		page.SaveTo(base)
	}
	for _, theme := range book.themes {
		theme.SaveTo(base)
	}
	return nil
}

func findH1(article *Element) string {
	var buf bytes.Buffer
	enc := NewHtmlEncoder(&buf)
	enc.Encode(article)
	from := bytes.Index(buf.Bytes(), []byte("<h1>")) + 4
	to := bytes.Index(buf.Bytes(), []byte("</h1>"))
	return strings.TrimSpace(string(buf.Bytes()[from:to]))
}

// AddPage creates a new page and returns a link to it
func (book *Book) AddPage(right string, article *Element) *Element {
	title := findH1(article)
	filename := filenameFrom(title) + ".html"

	page := newPage(
		filename,
		stripTags(title)+" - "+book.Title,
		pageHeader(right+" - "+A(Href("index.html"), book.Title).String()),
		article,
		Footer(book.Author),
	)
	book.pages = append(book.pages, page)
	return linkToPage(page)
}

func (me *Book) AddThemes(v ...*CSS) {
	me.themes = append(me.themes, v...)
}

func newPage(filename, title string, header, article, footer *Element) *Page {
	return NewFile(filename,
		Html(Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1.0"),
				),
				stylesheet("theme.css"),
				stylesheet("a4.css"),
				Title(title)),
			Body(
				header,
				article,
				footer,
			),
		),
	)
}

func linkToPage(page *Page) *Element {
	return Li(A(Href(page.Filename), findH1(page.Element)))
}

func stripTags(in string) string {
	var buf bytes.Buffer
	var inside bool
	for _, r := range in {
		switch r {
		case '<':
			inside = true
		case '>':
			inside = false
		default:
			if inside {
				continue
			}
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func filenameFrom(in string) string {
	tidy := bytes.NewBufferString("")
	var inside bool
	var last rune
	for _, c := range in {
		switch c {
		case '(', ')', '-':
			continue
		case ' ':
			if last == '_' {
				continue // skip two consecutive spaces
			}
			last = '_'
			tidy.WriteRune(last)
		case '<':
			inside = true
		case '>':
			inside = false
		default:
			if inside {
				continue
			}
			last = rune(strings.ToLower(string(c))[0])
			tidy.WriteRune(last)
		}
	}
	return tidy.String()
}

func pageHeader(right string) *Element {
	h := Header()
	if right != "" {
		h = h.With(Code(right))
	}
	return h
}

// stylesheet returns a link web element
func stylesheet(href string) *Element {
	return Link(Rel("stylesheet"), Type("text/css"), Href(href))
}

// Boxnote returns a small box aligned to the left with given top
// margin in cm.
func Sidenote(el interface{}, cm float64) *Element {
	return Div(Class("sidenote"),
		&Attribute{
			Name: "style",
			Val:  fmt.Sprintf("margin-top: %vcm", cm),
		},
		Div(Class("inner"), el),
	)
}

// loadFullFile returns a wrapped element with label and file contents.
// If label is empty string the filename last part is used.
func loadFullFile(label, filename string) *Element {
	if label == "" {
		dir := filepath.Base(filepath.Dir(filename))
		label = path.Join(dir, filepath.Base(filename))
	}
	return Wrap(
		Div(Class("filename"), label),
		loadFile(filename, 0, -1),
	)
}

// loadFile returns a pre web element wrapping the contents from the
// given file. If to == -1 all lines to the end of file are returned.
func loadFile(filename string, span ...int) *Element {
	from, to := 0, -1
	if len(span) == 2 {
		from, to = span[0], span[1]
	}
	v := files.MustLoadLines(filename, from, to)
	class := "srcfile"
	if from == 0 && to == -1 {
		class += " complete"
	}
	ext := filepath.Ext(filename)
	return Pre(Class(class), Code(Class(ext[1:]), v))
}

func gregoryv(name, txt string) *Element {
	return Li(
		fmt.Sprintf(
			`<a href="https://github.com/gregoryv/%s">%s</a> - %s`,
			name, name, txt,
		),
	)
}

// ShellCommand returns a web Element wrapping shell commands
func ShellCommand(v string) *Element {
	return Pre(Class("command"), Code(v))
}
