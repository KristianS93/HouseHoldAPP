package service



func (s *Server) ICreateMeal(w http.ResponseWriter, r *http.Request){
	meal.CreateMeal(&w, r)
}