package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/gregoryv/notes/page"
)

//go:generate go run .
func main() {
	fmt.Println("Dog")
	pages := map[string]writerTo{
		"dictionary.html": page.Dictionary,
		"index.html":      page.Index,
	}
	base := "./www/"
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
