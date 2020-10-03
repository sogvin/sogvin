package main

import (
	"fmt"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/fox"
	"github.com/gregoryv/wolf"
)

// NewStarCounter returns a new star counter command with logging
// enabled to cmd.Stderr.
func NewStarCounter(cmd wolf.Command) *StarCounter {
	return &StarCounter{
		Command: cmd,
		Logger:  fox.NewSyncLog(cmd.Stderr()).FilterEmpty(),
	}
}

type StarCounter struct {
	wolf.Command
	fox.Logger

	// star filters
	sizeFilter   string
	weightFilter string
}

// Run starts the application and waits for it to complete. Returns
// exit code 0 if completed ok, 1 otherwise.
func (me *StarCounter) Run() int {
	// Parse command line options
	cli := cmdline.New(me.Args()...)
	me.sizeFilter = cli.Option("-size").String("all")
	me.weightFilter = cli.Option("-weight").String("all")
	help := cli.Flag("-h, --help")

	if help {
		cli.WriteUsageTo(me.Stdout())
		return 0
	}
	if err := cli.Error(); err != nil {
		me.Log(err, ", try -help")
		return 1
	}
	if err := me.countStars(); err != nil {
		return 1
	}
	return 0
}

// countStars writes the result using the configured Stdout writer
func (me *StarCounter) countStars() error {
	// count stars using filters
	if me.weightFilter != "all" {
		return fmt.Errorf("bad weight")
	}
	return nil
}
