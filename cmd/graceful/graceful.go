package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gregoryv/logger"
)

func main() {
	me := NewMainEntry()
	me.Enter()
	os.Exit(me.Exit())
}

func NewMainEntry() *MainEntry {
	return &MainEntry{
		Logger: logger.NewProgress(),
	}
}

type MainEntry struct {
	logger.Logger
	err    error
	dryrun bool
}

func (me *MainEntry) Enter() {
	srv := NewServer(":2121")
	me.setupInterrupts(&srv)

	me.err = srv.ListenAndServe()
	// Wait for shutdown to complete
	<-srv.Done
}

func (me *MainEntry) setupInterrupts(srv *Server) {
	if me.skipf("interrupt with Ctrl-c") {
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)

	go func() {
		sig := <-c
		fmt.Printf("%v\n", sig)
		srv.Shutdown(context.Background())
	}()
}
