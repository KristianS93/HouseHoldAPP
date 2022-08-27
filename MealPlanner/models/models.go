package models

type Item struct {
	Id       int    `json:"Id"`
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

type Meal struct {
	Id          int    `json:"Id"`
	MealName    string `json:"MealName"`
	Description string `json:"Description"`
	Items       []Item `json:"Items"`
}

type WeekPlan struct {
	WeekNo      string `json:"WeekNo"`
	HouseHoldId string `json:"HouseHoldId"`
	Meals       []Meal `json:"Meals"`
}

type MealId struct {
	MealId int `json:"MealId"`
}

type HouseHold struct {
	HouseholdId   string `json:"HouseholdId"`
	GroceryListId string `json:"GroceryListId"`
}
