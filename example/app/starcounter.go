package main

import (
	"io"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/fox"
)

func main() {
	if err := NewStarCounter(os.Args...).Run(); err != nil {
		os.Exit(1)
	}
}

// Run starts the application and waits for it to complete.
func (me *StarCounter) Run() error {

	// Parse command line options
	cli := cmdline.New(me.Args...)
	me.sizeFilter = cli.Option("-size").String("all")
	me.weightFilter = cli.Option("-weight").String("all")
	help := cli.Flag("-h, -help")

	if help {
		cli.WriteUsageTo(me.Stdout)
		return nil
	}
	if err := cli.Error(); err != nil {
		me.Log(err, ", try -help")
		return err
	}
	return me.countStars()
}

// countStars writes the result using the configured Stdout writer
func (me *StarCounter) countStars() error {
	// count stars using filters
	return nil
}

// NewStarCounter returns an app, called from the operating system.
func NewStarCounter(args ...string) *StarCounter {
	wd, _ := os.Getwd()
	return &StarCounter{
		Env:    os.Environ(),
		Path:   args[0],
		Args:   args,
		Dir:    wd,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,

		Logger: fox.NewSyncLog(os.Stderr).FilterEmpty(),
	}
}

type StarCounter struct {
	Env    []string // environment variables
	Path   string   // path to executable that was invoked
	Args   []string // name and arguments
	Dir    string   // working directory
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	fox.Logger // for logging

	// star filters
	sizeFilter   string
	weightFilter string
}
