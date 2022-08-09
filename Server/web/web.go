package web

import (
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
	Templates map[string]*template.Template

	// Keeping track of current sessions by logging last activity on cookie key.
	Sessions map[string]time.Time
}
