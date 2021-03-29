package main

import (
	"encoding/json"
	"net/http"

	"github.com/gregoryv/sogvin/example/spaceflight"
)

func NewServer(srv *spaceflight.System) *Server {
	return &Server{srv: srv}
}

type Server struct {
	srv *spaceflight.System
}

func (me *Server) Router() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/routes", me.serveRoutes)
	return m
}

func (me *Server) serveRoutes(w http.ResponseWriter, r *http.Request) {
	var role spaceflight.Passenger
	me.srv.Use(&role)
	routes := role.ListRoutes()
	json.NewEncoder(w).Encode(routes)
}
