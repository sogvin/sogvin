package sogvin

import (
	"bytes"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

func NewWebsite() *Website {
	site := Website{
		Title:  "Software Engineering",
		Author: "Gregory Vin&ccaron;i&cacute;",
	}
	site.AddThemes(a4(), theme())

	toc := Article(Class("toc"),
		H1(site.Title),
		Img(Src("img/office.jpg")),
		P("Notes by ", site.Author),

		H2("Start"),
		Ul(
			site.AddPage("Start", gettingStartedWithProgramming()),
		),

		H2("Design"),
		Ul(
			site.AddPage("Design", purposeOfFuncMain()),
			site.AddPage("Design", nexusPattern()),
			site.AddPage("Design", gracefulServerShutdown()),
			site.AddPage("Design", componentsDiagram()),
			site.AddPage("Design", strictMode()),
			site.AddPage("Design", roleBasedService()),
		),
		H3("Go packages"),
		Ul(
			gregoryv("draw", "software engineering diagrams"),
			gregoryv("web", "html generation"),
		),

		H2("Test"),
		Ul(
			site.AddPage("Test", inlineTestHelpers()),
			site.AddPage("Test", alternateDesign()),
			site.AddPage("Test", setupTeardown()),
		),
		H3("Go packages"),
		Ul(
			gregoryv("golden", "simplify use of golden files"),
			gregoryv("qual", "quality constraints"),
			gregoryv("ex", "indented JSON or redirect handler response to stdout"),
			gregoryv("uncover", "paths that need more testing"),
		),

		H2("Build"),
		Ul(
			site.AddPage("Build", embedVersionAndRevision()),
		),
		H3("Go packages"),
		Ul(
			gregoryv("stamp", "build information code generator"),
			gregoryv("find", "files by name or content"),
		),
	)
	index := newPage("index.html", findH1(toc), Header(Code(
		versionField(), " - ", Released(),
	)), toc, Footer())
	site.pages = append(site.pages, index)

	return &site
}

type Website struct {
	Title  string
	Author string
	pages  []*Page
	themes []*CSS
}

// Saves all pages and table of contents
func (me *Website) SaveTo(base string) error {
	for _, page := range me.pages {
		page.SaveTo(base)
	}
	for _, theme := range me.themes {
		theme.SaveTo(base)
	}
	return nil
}

func (me *Website) AddThemes(v ...*CSS) {
	me.themes = append(me.themes, v...)
}

// ----------------------------------------

func versionField() *Element {
	el := Span()
	v := Version()
	if v == "unreleased" {
		el.With(Class("unreleased"), v)
	} else {
		el.With("v", v)
	}
	return el
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
func (me *Website) AddPage(right string, article *Element) *Element {
	title := findH1(article)
	filename := filenameFrom(title) + ".html"

	page := newPage(
		filename,
		stripTags(title)+" - "+me.Title,
		Header(Code(
			right+" - "+A(Href("index.html"), me.Title).String(),
		)),
		article,
		Footer(me.Author),
	)
	me.pages = append(me.pages, page)
	return linkToPage(page)
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

// stylesheet returns a link web element
func stylesheet(href string) *Element {
	return Link(Rel("stylesheet"), Type("text/css"), Href(href))
}

// Boxnote returns a small box aligned to the left with given top
// margin in cm.
func sidenote(el interface{}, cm float64) *Element {
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

func example(cmdline string, files ...string) *Element {
	res, err := runExample(cmdline, files...)
	if err != nil {
		log.Fatal(err)
	}
	return shellCommand(string(res))
}

// shellCommand returns a web Element wrapping shell commands
func shellCommand(v string) *Element {
	return Pre(Class("command"), Code(v))
}
