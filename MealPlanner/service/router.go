package service

//Routes handles the routing/handling of endpoints for the service
// The function is a method for the Server struct, and takes a pointer to a http.servemux type
func (s *Server) Routes() {
	s.Router.HandleFunc("/CreateMeal", s.ICreateMeal).Methods("GET")
}
