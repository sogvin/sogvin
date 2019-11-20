package main

import (
	"context"
	"net/http"
	"time"
)

func main() {
	me := &MainEntry{
		Server: http.Server{
			Addr:    ":2121",
			Handler: &http.ServeMux{},
		},
	}
	go me.ShutdownAfter(400 * time.Millisecond)
	me.Enter()
}

type MainEntry struct {
	err error
	http.Server
}

func (me *MainEntry) ShutdownAfter(d time.Duration) {
	<-time.After(d)
	me.Shutdown(context.Background())
}

func (me *MainEntry) Enter() {
	done := make(chan bool)
	me.RegisterOnShutdown(func() {
		// close everything before signaling done
		done <- true
	})

	me.err = me.ListenAndServe()
	<-done
}
