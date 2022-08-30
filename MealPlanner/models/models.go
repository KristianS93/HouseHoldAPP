package models

type Item struct {
	Id       int    `json:"Id"`
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

type Meal struct {
	Id          int64  `json:"Id"`
	MealName    string `json:"MealName"`
	Description string `json:"Description"`
	Items       []Item `json:"Items"`
}

type MealDB struct {
	Id          int64   `json:"Id"`
	MealName    string  `json:"MealName"`
	Description string  `json:"Description"`
	Items       []int64 `json:"Items"`
}

type Plan struct {
	Id          int64  `json:"Id"`
	WeekNo      int    `json:"WeekNo"`
	HouseHoldId string `json:"HouseHoldId"`
	Meals       []Meal `json:"Meals"`
}

type PlanId struct {
	PlanId int64 `json:"PlanId"`
}

type MealId struct {
	MealId int `json:"MealId"`
}

type HouseHold struct {
	Id            int    `json:"id"`
	HouseholdId   string `json:"HouseholdId"`
	GroceryListId string `json:"GroceryListId"`
	Plans         []Plan `json:"Plans"`
	Meals         []Meal `json:"Meals"`
}

type HouseHoldDB struct {
	Id            int     `json:"id"`
	HouseholdId   string  `json:"HouseholdId"`
	GroceryListId string  `json:"GroceryListId"`
	Plans         []int64 `json:"Plans"`
	Meals         []int64 `json:"Meals"`
}

type ItemIds struct {
	ItemIds []int64 `json:"ItemIds"`
}
