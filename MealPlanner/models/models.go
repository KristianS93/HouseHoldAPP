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

type MealDB struct {
	Id          int     `json:"Id"`
	MealName    string  `json:"MealName"`
	Description string  `json:"Description"`
	Items       []int64 `json:"Items"`
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
	Id            int    `json:"id"`
	HouseholdId   string `json:"HouseholdId"`
	Meals         string `json:"Meals"`
	GroceryListId string `json:"GroceryListId"`
}

type ItemIds struct {
	ItemIds []int64 `json:"ItemIds"`
}
