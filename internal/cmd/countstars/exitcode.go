package main

import (
	"fmt"
	"io"
	"os"
)

type ExitCode = int

const (
	ExitOk ExitCode = iota
	ExitFail
)

// CountGalaxies writes the number of galaxies in the universe to the
// given writer.
func CountGalaxies(w io.Writer) ExitCode {
	universe, err := listGalaxiesInUniverse()
	if err != nil {
		return ExitFail
	}
	fmt.Fprint(w, len(universe))
	return ExitOk
}

func listGalaxiesInUniverse() ([]byte, error) {
	return nil, nil
}

func tryexit() {
	os.Exit(CountGalaxies(os.Stdout))
}
