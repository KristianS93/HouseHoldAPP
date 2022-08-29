package service

import (
	"mealplanner/service/household"
	"mealplanner/service/item"
	"mealplanner/service/meal"
	"mealplanner/service/plan"
	"net/http"
)

// Item functionalites
func (s *Server) IAddItem(w http.ResponseWriter, r *http.Request) {
	item.AddItem(w, r)
}
func (s *Server) IDeleteItem(w http.ResponseWriter, r *http.Request) {
	item.DeleteItem(w, r)
}
func (s *Server) IUpdateItem(w http.ResponseWriter, r *http.Request) {
	item.UpdateItem(w, r)
}

// Meal functionalities
func (s *Server) ICreateMeal(w http.ResponseWriter, r *http.Request) {
	meal.CreateMeal(w, r)
}
func (s *Server) IGetMeal(w http.ResponseWriter, r *http.Request) {
	meal.GetMeal(w, r)
}
func (s *Server) IChangeMeal(w http.ResponseWriter, r *http.Request) {
	meal.ChangeMeal(w, r)
}
func (s *Server) IDeleteMeal(w http.ResponseWriter, r *http.Request) {
	meal.DeleteMeal(w, r)
}

// Plan functionalities
func (s *Server) ICreatePlan(w http.ResponseWriter, r *http.Request) {
	plan.CreatePlan(w, r)
}
func (s *Server) IGetPlan(w http.ResponseWriter, r *http.Request) {
	plan.GetPlan(w, r)
}
func (s *Server) IChangePlan(w http.ResponseWriter, r *http.Request) {
	plan.ChangePlan(w, r)
}
func (s *Server) IGeneratePlan(w http.ResponseWriter, r *http.Request) {
	plan.GeneratePlan(w, r)
}
func (s *Server) IGenerateList(w http.ResponseWriter, r *http.Request) {
	plan.GenerateList(w, r)
}
func (s *Server) IDeletePlan(w http.ResponseWriter, r *http.Request) {
	plan.DeletePlan(w, r)
}

// Household functionalities
func (s *Server) ICreateHousehold(w http.ResponseWriter, r *http.Request) {
	household.CreateHouseHold(w, r)
}
func (s *Server) IDeleteHousehold(w http.ResponseWriter, r *http.Request) {
	household.DeleteHouseHold(w, r)
}
func (s *Server) ICreateGroceryList(w http.ResponseWriter, r *http.Request) {
	household.CreateGroceryList(w, r)
}
