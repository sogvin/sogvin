package sogvin

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

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

func example(args string, files ...string) *Element {
	res, err := runExample(args, files...)
	if err != nil {
		log.Fatal(err)
	}
	return shellCommand(string(res))
}

func loadExample(filename string) *Element {
	src := loadAs(filename, "init", "main")
	i := strings.Index(src, "\n") // first line
	fn := strings.Index(src, "\npackage")
	var block string
	if fn > i {
		block = src[i+1 : fn]
		block = strings.ReplaceAll(block, "//", "")
	}
	e := Wrap(
		H1(src[3:i]),
		P(block),
		Div(Class("filename"), filename),
		Pre(Class("srcfile complete"),
			Code(Class("go"), src[fn+1:]),
		),
	)
	return e
}

func skipFirstLine(in string) string {
	i := strings.Index(in, "\n") // skip first line, which should be comment
	return in[i+1:]
}

func loadAs(filename, fn, as string) string {
	data := files.MustLoad(filename)
	return strings.ReplaceAll(data, fn, as)
}

// shellCommand returns a web Element wrapping shell commands
func shellCommand(v string) *Element {
	return Pre(Class("command"), Code(v))
}

func linkDrill(filename string) *Element {
	title := drillTitle(filename)
	return Li(A(Href(toHtmlFile(filename)), title))
}

func drillTitle(filename string) string {
	line := firstLine(filename)
	title := line
	parts := strings.Split(line, ";")
	if len(parts) > 1 {
		title = parts[1]
	}
	return title
}

func toHtmlFile(filename string) string {
	return strings.Replace(filename, ".go", ".html", 1)
}

func firstLine(filename string) string {
	fh := openFile(filename)
	defer fh.Close()
	line := readLine(fh)
	if line[:2] != "//" {
		log.Fatal("missing file comment: ", filename)
	}
	return line[3:] // skip first comment '// '
}

func openFile(filename string) io.ReadCloser {
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fh
}

func readLine(r io.Reader) string {
	s := bufio.NewScanner(r)
	s.Scan()
	return s.Text()
}

// runExample first file contains init() that is renamed to main()
func runExample(args string, files ...string) ([]byte, error) {
	first := files[0]

	data, err := ioutil.ReadFile(first)
	if err != nil {
		return nil, err
	}

	data = bytes.ReplaceAll(data, []byte("func init("), []byte("func main("))
	data = bytes.ReplaceAll(data, []byte("package drill"), []byte("package main"))

	// Use name of first file as command name, so we can have many
	// files in same directory, but speed up builds
	name := filepath.Base(first)
	i := strings.Index(name, ".")
	dir := filepath.Join(os.TempDir(), name[:i])
	if err := os.MkdirAll(dir, 0722); err != nil {
		return nil, err
	}
	log.Println(dir)

	scriptFile := filepath.Join(dir, name)
	if err := ioutil.WriteFile(scriptFile, data, 0644); err != nil {
		return nil, err
	}

	parts := strings.Split(args, " ")
	fullArgs := append(
		[]string{"run", filepath.Base(scriptFile)}, parts...,
	)

	cmd := exec.Command("go", fullArgs...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()

	var buf bytes.Buffer
	buf.WriteString("$ ")
	cmdline := cmd.String()
	cmdline = strings.Replace(cmdline, "/home/gregory/dl/go1/go/bin/", "", 1)
	buf.WriteString(cmdline)
	buf.WriteString("\n")
	buf.Write(out)
	return buf.Bytes(), err
}
