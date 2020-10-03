package main

import (
	"os"

	"github.com/gregoryv/wolf"
)

func main() {
	cmd := wolf.NewOSCmd()
	sc := NewStarCounter(cmd)
	exitCode := sc.Run()
	os.Exit(exitCode)
}
