package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	done := make(chan bool)
	graceful := func() {
		// close everything before signaling done
		fmt.Println("Closing down ...")
		done <- true
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: &http.ServeMux{},
	}
	srv.RegisterOnShutdown(graceful)
	go stopOn(&srv, os.Kill, os.Interrupt)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	// Wait for graceful to complete
	<-done
}

// stopOn calls Shutdown on the server for the given signals
func stopOn(srv *http.Server, signals ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	sig := <-c
	fmt.Printf("%v\n", sig)
	srv.Shutdown(context.Background())
}
