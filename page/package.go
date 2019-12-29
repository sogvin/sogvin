package page

import (
	"fmt"
	"io"
	"os"
	"path"
)

func WriteAllPages(base string) {
	pages := map[string]writerTo{
		"dictionary.html":    Dictionary,
		"index.html":         Index,
		"nexus_pattern.html": NexusPattern,
	}
	for filename, page := range pages {
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
}

type writerTo interface {
	WriteTo(io.Writer) (int, error)
}
