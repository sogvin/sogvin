package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gregoryv/notes/page"
)

//go:generate go run .
func main() {
	fmt.Println("Dog")
	pages := map[string]writerTo{
		"../dictionary.html": page.Dictionary,
	}
	for filename, page := range pages {
		fmt.Println("  ", filename)
		fh, err := os.Create(filename)
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
