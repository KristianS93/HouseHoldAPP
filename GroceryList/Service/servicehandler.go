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

	//The pointer has to be zero when, running the program first time
	//And we can create the Multiplexer so we can route.
	// if s.Router == nil {
	// 	s.Router = http.NewServeMux()

	// 	//Getting the routes for this service
	// 	s.Routes(s.Router)
	// }

	//Setting the specified host and port
	s.HostName = ConstHost
	s.HostPort = ConstPort

	s.Router = mux.NewRouter()
	s.Routes()
}

//Run creates the listen and serve, which turns on the server, and then awaits requests.
func (s *Server) Run(addr string) {
	// fmt.Println("Service starting up at: http://" + s.HostName + s.HostPort)
	// err := http.ListenAndServe((s.HostName + s.HostPort), s.Router)
	// if err != nil {
	// 	log.Fatalln("Failed to start GroceryList service, exiting.")
	// }

	log.Fatal(http.ListenAndServe(addr, s.Router))

}

// EnableCors take the pointer to the ResponseWriter and sets the CORS
// Input has to be an ADDRESS to a responsewriter.
func EnableCors(w *http.ResponseWriter) {
	//Setting the cors for the pointer to the responsewriter.
	(*w).Header().Set("Access-Control-Allow-Origin", ConstAllowedCORS)
}
