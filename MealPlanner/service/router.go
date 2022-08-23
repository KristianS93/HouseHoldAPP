package service

//Routes handles the routing/handling of endpoints for the service
// The function is a method for the Server struct, and takes a pointer to a http.servemux type
func (s *Server) Routes() {
	//Meal endpoints
	s.Router.HandleFunc("/CreateMeal", s.ICreateMeal).Methods("POST")
	s.Router.HandleFunc("/GetMeal", s.IGetMeal).Methods("GET")
	s.Router.HandleFunc("/ChangeMeal", s.IChangeMeal).Methods("PATCH")
	s.Router.HandleFunc("/DeleteMeal", s.ICreateMeal).Methods("DELETE")
	
	//Plan endpoints
	s.Router.HandleFunc("/CreatePlan", s.ICreatePlan).Methods("POST")
	s.Router.HandleFunc("/GetPlan", s.IGetPlan).Methods("GET")
	s.Router.HandleFunc("/ChangePlan", s.IChangePlan).Methods("PATCH")
	s.Router.HandleFunc("/GeneratePlan", s.IGeneratePlan).Methods("POST") //måske get med url param
	s.Router.HandleFunc("/GenerateList", s.IGenerateList).Methods("POST")
	s.Router.HandleFunc("/DeletePlan", s.IDeletePlan).Methods("DELETE")
}
