package main

import (
	"io"

	"github.com/gregoryv/logger"
)

var (
	debug = logger.Silent
)

// CountStars writes the number of starts in the given galaxy to the
// given writer.
func CountStars(w io.Writer, galaxy string) {
	debug.Logf("counting stars in %s", galaxy)
	// ...
}

func SetVerbose(yes bool) {
	if !yes {
		return
	}
	debug = logger.New()
}
