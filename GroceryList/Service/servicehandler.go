package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const SessionTimeOut = ConstSessionTimeOut

type Server struct {
	//Router will be a pointer to the http request multiplexer
	Router *mux.Router

	//Setting the host
	HostName string
	//Setting the host port
	HostPort string

	//Loggin the current session by logging the last activity from a cookie
	Sessions map[string]time.Time
}

// Init initializing MUX settings
func (s *Server) Init() {

	//Setting the specified host and port
	s.HostName = ConstHost
	s.HostPort = ConstPort

	s.Router = mux.NewRouter()
	s.Routes()
}

// Run creates the listen and serve, which turns on the server, and then awaits requests.
func (s *Server) Run(addr string) {
	log.Println("Service starting up at: http://" + s.HostName + s.HostPort)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
