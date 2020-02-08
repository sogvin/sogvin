package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version  = "0.0"
	revision = "dev"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "print version and exit")

	if showVersion {
		fmt.Printf("%s", version)
		os.Exit(0)
	}
}
