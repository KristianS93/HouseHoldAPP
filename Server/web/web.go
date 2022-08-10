package web

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

const (
	// The amount of seconds a session should stay active, before a new login is required
	// to continue utilizing the service.
	SessionTimeOut = 1200
)

type Server struct {
	// Will serve as the router used to route incoming requests properly.
	Router *http.ServeMux

	// Domain and port, locally would be "localhost" and ":8080", or similar ports.
	HostName string
	HostPort string

	// Storing all current templates, to be ready for execution.
	Templates *template.Template

	// Keeping track of current sessions by logging last activity on cookie key.
	Sessions map[string]time.Time
}

func (s *Server) Init() {
	if s.Router == nil {
		// This ensures that incoming traffic reaches the designated router.
		s.Router = http.NewServeMux()
		s.Routes(s.Router)
	}

	// Need to implement an init for
	// s.Templates

	if s.Sessions == nil {
		s.Sessions = make(map[string]time.Time)
	}

	if s.HostName == "" {
		s.HostName = "localhost"
		log.Println("No HostName specified, defaulting to localhost")
	}
	if s.HostPort == "" {
		s.HostPort = ":8888"
		log.Println("No HostPort specified, defaulting to :8888")
	}
}

// Run launches a LAS on the specified HostName and Port,
// while using the Server.Router as the ServeMux.
func (s *Server) Run() {
	err := http.ListenAndServe((s.HostName + s.HostPort), s.Router)
	if err != nil {
		log.Fatalln("Failed to start a server, closing application.")
	}
}



