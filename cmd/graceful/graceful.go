package main

import (
	"context"
	"fmt"
	"net/http"
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
		Server: http.Server{
			Addr:    ":2121",
			Handler: &http.ServeMux{},
		},
	}
}

type MainEntry struct {
	logger.Logger
	err    error
	dryrun bool

	http.Server
}

func (me *MainEntry) Enter() {
	me.setupInterrupts()
	done := make(chan bool)
	me.RegisterOnShutdown(func() {
		// close everything before signaling done
		done <- true
	})

	me.err = me.ListenAndServe()
	<-done
}

func (me *MainEntry) setupInterrupts() {
	if me.skipf("interrupt with Ctrl-c") {
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)

	go func() {
		<-c
		fmt.Println("exiting...")
		me.Shutdown(context.Background())
	}()
}
