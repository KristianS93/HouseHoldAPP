package service

import (
	"net/http"
)

// Routes handles the routing/handling of endpoints for the service
// The function is a method for the Server struct, and takes a pointer to a http.servemux type
func (s *Server) Routes(r *http.ServeMux) {
	r.HandleFunc("/AddItem", s.AddItem)
	r.HandleFunc("/DeleteItem", s.DeleteItem)
	r.HandleFunc("/ChangeItem", s.ChangeItem)

	r.HandleFunc("/GetList", s.GetList)
	r.HandleFunc("/CreateList", s.CreateList)
	r.HandleFunc("/DeleteList", s.DeleteList)
	r.HandleFunc("/ClearList", s.ClearList)
}
