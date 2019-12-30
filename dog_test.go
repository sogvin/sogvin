package notes

import (
	"fmt"
	"io"
	"os"
	"path"
	"testing"
)

func Test_generate_www(t *testing.T) {
	WriteAllPages("./se")
}

func WriteAllPages(base string) {
	pages := map[string]writerTo{
		"dictionary.html":               Dictionary,
		"graceful_server_shutdown.html": GracefulServerShutdown,
		"index.html":                    Index,
		"inline_test_helpers.html":      InlineTestHelpers,
		"nexus_pattern.html":            NexusPattern,
		"purpose_of_func_main.html":     PurposeOfFuncMain,
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
