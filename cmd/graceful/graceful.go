package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{Addr: ":2121", Handler: &http.ServeMux{}}

	done := make(chan bool)
	srv.RegisterOnShutdown(func() {
		fmt.Println("Stopping")
		time.Sleep(100 * time.Millisecond)
		done <- true
	})

	go shutdownAfter(srv, 400*time.Millisecond)

	srv.ListenAndServe()
	fmt.Println("Waiting for done")
	<-done
}

func shutdownAfter(srv *http.Server, d time.Duration) {
	<-time.After(d)
	srv.Shutdown(context.Background())
}
