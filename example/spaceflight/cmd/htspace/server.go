package main

import (
	"encoding/json"
	"net/http"

	"github.com/gregoryv/sogvin/example/spaceflight"
)

func NewServer(sys *spaceflight.System) *Server {
	return &Server{sys: sys}
}

type Server struct {
	sys *spaceflight.System
}

func (me *Server) Router() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/routes", me.serveRoutes)
	return m
}

func (me *Server) serveRoutes(w http.ResponseWriter, r *http.Request) {
	// Default to the passenger role
	var role spaceflight.Passenger
	me.sys.Use(&role)

	routes := role.ListRoutes()
	json.NewEncoder(w).Encode(routes)
}
