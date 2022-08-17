package service

// Routes handles the routing/handling of endpoints for the service
// The function is a method for the Server struct, and takes a pointer to a http.servemux type
func (s *Server) Routes() {
	s.Router.HandleFunc("/AddItem", s.AddItem).Methods("POST")
	s.Router.HandleFunc("/DeleteItem", s.DeleteItem).Methods("DELETE")
	s.Router.HandleFunc("/ChangeItem", s.ChangeItem).Methods("PATCH")

	s.Router.HandleFunc("/GetList", s.GetList).Methods("GET")
	s.Router.HandleFunc("/CreateList", s.CreateList).Methods("POST")
	s.Router.HandleFunc("/DeleteList", s.DeleteList).Methods("DELETE")
	s.Router.HandleFunc("/ClearList", s.ClearList).Methods("DELETE")
}
