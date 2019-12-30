package notes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	. "github.com/gregoryv/web/doctype"
)

func Test_generate_www(t *testing.T) {
	WriteAllPages("./se")
}

func WriteAllPages(base string) {
	toc := Ul(Class("toc"))
	pages := map[string]writerTo{
		"purpose_of_func_main.html":     page(PurposeOfFuncMain, "func main()"),
		"nexus_pattern.html":            page(NexusPattern, "Nexus pattern"),
		"inline_test_helpers.html":      page(InlineTestHelpers, "Testing"),
		"graceful_server_shutdown.html": page(GracefulServerShutdown, "Shutdown"),

		//		"dictionary.html": page(Dictionary, "Dictionary"),
	}
	for filename, p := range pages {
		saveAs(base, filename, p)
		toc = toc.With(
			Li(
				A(
					Href(filename),
					findH1(p),
				),
			),
		)
	}
	index := Article(
		H1("Software Engineering"),
		P("Notes by", myname),

		H2("Table of Contents"),
		toc,
	)
	saveAs(base, "index.html", page(index, ""))
}

func findH1(article writerTo) string {
	var buf bytes.Buffer
	article.WriteTo(&buf)
	from := bytes.Index(buf.Bytes(), []byte("<h1>")) + 4
	to := bytes.Index(buf.Bytes(), []byte("</h1>"))
	return strings.TrimSpace(string(buf.Bytes()[from:to]))
}

func saveAs(base, filename string, page writerTo) {
	out := path.Join(base, filename)
	fmt.Println("  ", out)
	fh, err := os.Create(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	page.WriteTo(fh)
	fh.Close()
}

type writerTo interface {
	WriteTo(io.Writer) (int, error)
}

func page(article *Tag, right string) *HtmlTag {
	return Html(en,
		Head(utf8, viewport, theme, a4),
		Body(
			header("", right),
			article,
			footer,
		),
	)
}
