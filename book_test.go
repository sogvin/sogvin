package notes

import (
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	. "github.com/gregoryv/web/doctype"
)

func Test_generate_www(t *testing.T) {
	WriteAllPages("./se")
}

func WriteAllPages(base string) {
	pages := map[string]writerTo{
		"dictionary.html":               page(Dictionary, "Dictionary"),
		"graceful_server_shutdown.html": page(GracefulServerShutdown, "Shutdown"),
		"index.html":                    page(Index, ""),
		"inline_test_helpers.html":      page(InlineTestHelpers, "Testing"),
		"nexus_pattern.html":            page(NexusPattern, "Nexus pattern"),
		"purpose_of_func_main.html":     page(PurposeOfFuncMain, "func main()"),
	}
	for filename, art := range pages {
		out := path.Join(base, filename)
		fmt.Println("  ", out)
		fh, err := os.Create(out)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		art.WriteTo(fh)
		fh.Close()
	}
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
