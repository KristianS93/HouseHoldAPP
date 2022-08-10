package service

import (
	"fmt"
	"net/http"
)

// Routes handles the routing/handling of endpoints for the service
// The function is a method for the Server struct, and takes a pointer to a http.servemux type
func (s *Server) Routes(r *http.ServeMux) {
	r.HandleFunc("/GetList", s.GetList)
}

func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {
	// Handle cors
	EnableCors(&w)
	fmt.Println("Can get")
	//Check for correct method

	//correct logic to get the specified list

	//Serve the json for the list: json object with an array of items

}
