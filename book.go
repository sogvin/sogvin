package notes

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gregoryv/notes/internal"
	. "github.com/gregoryv/web"
)

func NewBook() *Book {
	book := &Book{}

	toc := Article(
		H1("Software Engineering"),
		P("Notes by ", myname),
		H2("Table of Contents"),

		H3("Design"),
		Ul(Class("toc"),
			book.AddPage("func main()", PurposeOfFuncMain),
			book.AddPage("Nexus pattern", NexusPattern),
			book.AddPage("Shutdown", GracefulServerShutdown),

			gregoryv("draw", "software engineering diagrams"),
			gregoryv("web", "html generation"),
		),

		H3("Test"),
		Ul(Class("toc"),
			book.AddPage("Testing", InlineTestHelpers),
			gregoryv("golden", "simplify use of golden files"),
			gregoryv("qual", "quality constraints"),
			gregoryv("ex", "indented JSON or redirect handler response to stdout"),
			gregoryv("uncover", "paths that need more testing"),
		),

		H3("Build"),
		Ul(Class("toc"),
			gregoryv("stamp", "build information code generator"),
			gregoryv("find", "files by name or content"),
		),
	)
	index := newPage("index.html", "", toc)
	book.pages = append(book.pages, index)
	return book
}

func linkToPage(page *Page) *Element {
	return Li(A(Href(page.Filename), findH1(page.Element)))
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

var GracefulServerShutdown = Article(
	H1("Graceful server shutdown"),
	P(`Avoid disrupting ongoing requests by shutting down
	   gracefully. In the below example Ctrl-c can be used to signal
	   an interrupt which tells a listening <code>http.Server</code>
	   to shutdown.`),

	boxnote("Register the graceful part of the server.", 4.8),
	boxnote("Important to wait for graceful stop to end.", 7.8),
	loadGoFile("./internal/cmd/graceful/graceful.go", 11, -1),
	P(`Remember that you could expose the Shutdown func of your
       server through an URL to simplify clean shutdown. Useful for
       when you are doing continuous integration and
       deployment.`),
)

var Dictionary = Article(
	H1("Dictionary"),
	P(`Short list of words/terms often used in software engineering
	   and sometimes defined differently. Only domain agnostic terms
	   have been listen, for the rest consult an english dictionary.
	   I often use the <code>dict</code> command line tool.`),

	Dl(
		Dt("Argument"),
		Dd("String following the command on the command line."),

		Dt("Flag"),
		Dd("Boolean option."),

		Dt("Option"),
		Dd("Argument starting with single or double dashes."),
	),
)

var InlineTestHelpers = Article(
	H1("Inline test helpers"),
	P(`Use inline test helpers to minimize indentation and have
       failures point out failed cases directly.`),

	boxnote("Inlined helper does not need t argument.", 0.8),
	boxnote("Descriptive cases fail on correct line.", 5.6),
	loadGoFile("./internal/testing/inline_test.go", 8, -1),
	boxnote("Utmost 2 inlined helpers.", 0.2),

	P(`Keep it simple and use utmost two inlined helpers. Compared to
       table-driven-tests inlined helpers declare the <em>how</em>
       before the cases.  If you have many cases, this style is more
       readable as you first tell the reader the meaning of
       &#34;ok&#34; and &#34;bad&#34;.  <br> Another positive benefit
       of this style is values are not grouped in a testcase
       variable. I.e. readability improves as the values are used
       directly.  <br>This style may be less readable if each case
       requires many values, though it depends on the lenght of the
       values combined.`),
)

var NexusPattern = Article(
	H1("Nexus pattern"),
	P(
		"The word nexus is defined as",
		Quote(
			"&#34;The means of connection between things linked in series&#34;",
		),
		"The pattern is useful in",
		A(
			Href(
				"https://go.googlesource.com/proposal/+/master/design"+
					"/go2draft-error-handling-overview.md",
			),
			"error handling",
		),
		"sequential function calls.",
	),

	H2("Example <code>CopyFile(from, to string)</code>"),
	P(`Copying a file, if done all in one function, is unreadable due
       to multiple error checking and handling.  With the nexus
       pattern you define a <code>type fileIO struct</code> with the
       error field. Each method must check the previous error and
       return if it is set without doing anything. This way all
       subsequent calls are no-operations.`),

	boxnote("The err field links operations.", 0.6),
	boxnote("Each method sets x.err before returning.", 3.3),
	loadGoFile("./internal/errhandling/nexus.go", 21, -1),

	`With the fileIO nexus inplace the CopyFile function is
	readable and with only one error checking and handling needed.`,
	loadGoFile("./internal/errhandling/nexus.go", 8, 19),
)

var PurposeOfFuncMain = Article(
	H1("Purpose of <code>func main()</code>"),

	P(`The purpose of <code>func main()</code> is to <b>translate
       commandline arguments to application startup state</b>. Once
       the state is prepared a specific entry function is called. More
       often than not, logging verbosity is one such state that needs
       to be configured early on.  <br> Use the builtin flag package
       to define, document and parse the arguments.`),

	H2("Example <code>CountStars(galaxy string)</code>"),

	P(`Imagine an application that counts the stars in a named
	   galaxy. The main function should then make sure the flags are
	   correct and forward them as arguments to the function doing the
	   actual work. The name of the galaxy would be such a flag and
	   perhaps a verbosity flag for debugging purposes.`),

	loadGoFile("./internal/cmd/countstars/main.go", 8, -1),

	P(`Now that you know what the main function should do, let us take
	   a look at how it should be done, apart of the flag definition
	   and argument passing.<br> First, the cyclomatic complexity of
	   the main function is one. Ie. there is only one path through
	   this program.  There are however two exit points, apart from
	   the obvious one <code>flag.Parse()</code> exits if the parsed
	   flags do not match the predefined. The single pathway means
	   that testing the main function is simple. Execute this
	   application with valid flags and all lines are covered, leaving
	   all other code for unittesting.<br> Also, if you execute the
	   program you would note that second, the order of the flags are
	   sorted in the same way as the help output.`),

	boxnote("Cyclomatic complexity should be one.", -5.2),
	boxnote("Flag order should match output.", -1.7),

	H2("Benefits"),

	P(`Adhering to the &ldquo;keep it simple principle&rdquo; and only
	   doing one thing in each function, works out nicely for the main
	   function as well. One could argue that, if you moved everything
	   inside main into a start function, the flag definitions would
	   also be tested.  Think about it for a minute and figure out
	   what exactly you would be testing. If the flag package already
	   makes sure it's functions work as expected the only thing left
	   is testing which flags you have defined.  They would need to
	   be updated each time you add or remove a flag which is a sign
	   of a poor test.<br> You could potentially refactor main and
	   separate the flag definitions into smaller functions for
	   readability but you still wouldn't need to write unittests for
	   them.`),

	P(`Keep main simple, constrain it to only set global startup state
	   before calling the one function that does the actual
	   work.<br>This works great for services and simpler commands
	   that only do one thing.`),
)

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
		right+" - "+A(Href("index.html"), "Software Engineering").String(),
		article,
	)
	book.pages = append(book.pages, page)
	return linkToPage(page)
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

func newPage(filename, right string, article *Element) *Page {
	return NewPage(filename,
		Html(en,
			Head(utf8, viewport, theme, a4),
			Body(header(right), article, footer),
		),
	)
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

func loadGoFile(filename string, from, to int) string {
	return `<pre class="srcfile"><code class="go">` +
		internal.LoadFile(filename, from, to) + "</code></pre>"
}

func gregoryv(name, txt string) *Element {
	return Li(
		fmt.Sprintf(
			`<a href="https://github.com/gregoryv/%s">%s</a> - %s`,
			name, name, txt,
		),
	)
}
