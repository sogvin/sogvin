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

	size   string
	weight string
}

// Run starts the application and waits for it to complete. Returns
// exit code 0 if completed ok, 1 otherwise.
func (me *StarCounter) Run() int {
	// Parse command line options
	var (
		cli    = cmdline.NewParser(me.Args()...)
		size   = cli.Option("-size").String("all")
		weight = cli.Option("-weight").String("all")
		help   = cli.Flag("-h, --help")
	)

	switch {
	case help:
		cli.WriteUsageTo(me.Stdout())
		return 0

	case !cli.Ok():
		me.Log(cli.Error(), ", try --help")
		return 1
	}
	me.size = size
	me.weight = weight
	if err := me.countStars(); err != nil {
		return 1
	}
	return 0
}

// countStars writes the result using the configured Stdout writer
func (me *StarCounter) countStars() error {
	// count stars using filters
	if me.weight != "all" {
		return fmt.Errorf("bad weight")
	}
	return nil
}
