package main

import (
	"flag"
	"os"
)

func main() {
	galaxy := "milkyway"
	flag.StringVar(
		&galaxy, "galaxy",
		galaxy, "name of galaxy to count stars in",
	)
	verbose := false
	flag.BoolVar(
		&verbose, "verbose",
		verbose, "enables verbose logging",
	)
	flag.Parse()
	SetVerbose(verbose)
	CountStars(os.Stdout, galaxy)
}
