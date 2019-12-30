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
		"dictionary.html":          Dictionary,
		"index.html":               Index,
		"nexus_pattern.html":       NexusPattern,
		"inline_test_helpers.html": InlineTestHelpers,
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
