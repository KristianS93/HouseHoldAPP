package service

import "github.com/gorilla/mux"

type Server struct {
	Router   *mux.Router
	HostName string
	HostPort string

	// db connection
	DB string
}
