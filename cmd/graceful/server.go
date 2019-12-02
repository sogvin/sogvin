package main

import (
	"net/http"

	"github.com/gregoryv/logger"
)

func NewServer(bind string) Server {
	srv := Server{
		Logger: logger.Silent,
		Server: http.Server{
			Addr:    bind,
			Handler: &http.ServeMux{},
		},
		Done: make(chan bool),
	}
	srv.RegisterOnShutdown(func() {
		srv.Log("Shutting down...")
		// close everything before signaling done
		srv.Done <- true
	})
	return srv
}

type Server struct {
	logger.Logger
	http.Server

	Done chan bool
}
