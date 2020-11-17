package main

import (
	"os"

	"github.com/gregoryv/wolf"
)

func main() {
	var (
		cmd  = wolf.NewOSCmd()
		sc   = NewStarCounter(cmd)
		code = sc.Run()
	)
	os.Exit(code)
}
