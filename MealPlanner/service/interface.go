package service

import (
	"mealplanner/service/meal"
	"mealplanner/service/plan"
	"net/http"
)

// Meal functionalities
func (s *Server) ICreateMeal(w http.ResponseWriter, r *http.Request) {
	meal.CreateMeal(&w, r)
}
func (s *Server) IGetMeal(w http.ResponseWriter, r *http.Request) {
	meal.GetMeal(&w, r)
}
func (s *Server) IChangeMeal(w http.ResponseWriter, r *http.Request) {
	meal.ChangeMeal(&w, r)
}
func (s *Server) IDeleteMeal(w http.ResponseWriter, r *http.Request) {
	meal.DeleteMeal(&w, r)
}

// Plan functionalities
func (s *Server) ICreatePlan(w http.ResponseWriter, r *http.Request) {
	plan.CreatePlan(&w, r)
}
func (s *Server) IGetPlan(w http.ResponseWriter, r *http.Request) {
	plan.GetPlan(&w, r)
}
func (s *Server) IChangePlan(w http.ResponseWriter, r *http.Request) {
	plan.ChangePlan(&w, r)
}
func (s *Server) IGeneratePlan(w http.ResponseWriter, r *http.Request) {
	plan.GeneratePlan(&w, r)
}
func (s *Server) IGenerateList(w http.ResponseWriter, r *http.Request) {
	plan.GenerateList(&w, r)
}
func (s *Server) IDeletePlan(w http.ResponseWriter, r *http.Request) {
	plan.DeletePlan(&w, r)
}
