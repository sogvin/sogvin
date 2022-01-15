package main

import (
	"github.com/gregoryv/cmdline"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		size   = cli.Option("-size").String("all")
		weight = cli.Option("-weight").Int(0)
	)
	cli.Parse()

	sc := NewStarCounter()
	sc.SetSize(size)
	sc.SetWeight(weight)

	if err := sc.Run(); err != nil {
		cmdline.DefaultShell.Fatal(err)
	}
}
