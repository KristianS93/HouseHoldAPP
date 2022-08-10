package service

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const SessionTimeOut = ConstSessionTimeOut

type Server struct {
	//Router will be a pointer to the http request multiplexer
	Router *http.ServeMux

	//Setting the host
	HostName string
	//Setting the host port
	HostPort string

	//Loggin the current session by logging the last activity from a cookie
	Sessions map[string]time.Time
}

//Init initializing MUX settings
func (s *Server) Init() {
	
	//The pointer has to be zero when, running the program first time
	//And we can create the Multiplexer so we can route.
	if s.Router == nil {
		s.Router = http.NewServeMux()

		//Getting the routes for this service
		s.Routes(s.Router)
	}

	//Setting the specified host and port
	s.HostName = ConstHost
	s.HostPort = ConstPort
}

func (s *Server) Run() {
	fmt.Println("Service starting up at: http://" + s.HostName + s.HostPort)
	err := http.ListenAndServe((s.HostName + s.HostPort), s.Router)
	if err != nil {
		log.Fatalln("Failed to start GroceryList service, exiting.")
	}
}

//EnableCors take the pointer to the ResponseWriter and sets the CORS
// Input has to be an ADDRESS to a responsewriter.
func EnableCors(w *http.ResponseWriter) {
	//Setting the cors for the pointer to the responsewriter.
	(*w).Header().Set("Access-Control-Allow-Origin", ConstAllowedCORS)
}


