package models

type Item struct {
	Id       int    `json:"Id"`
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

type ItemRequest struct {
	ListId   string `json:"ListId"`
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

type Meal struct {
	Id          int64  `json:"Id"`
	HouseholdId string `json:"HouseHoldId"`
	MealName    string `json:"MealName"`
	Description string `json:"Description"`
	Items       []Item `json:"Items"`
}

type MealDB struct {
	Id          int64   `json:"Id"`
	MealName    string  `json:"MealName"`
	HouseholdId string  `json:"HouseHoldId"`
	Description string  `json:"Description"`
	Items       []int64 `json:"Items"`
}

type Plan struct {
	Id          int64  `json:"Id"`
	WeekNo      int    `json:"WeekNo"`
	HouseHoldId string `json:"HouseHoldId"`
	Meals       []Meal `json:"Meals"`
}

type PlanDB struct {
	Id          int64   `json:"Id"`
	WeekNo      int     `json:"WeekNo"`
	HouseHoldId string  `json:"HouseHoldId"`
	Meals       []int64 `json:"Meals"`
}

type ReturnPlan struct {
	Id          int64  `json:"Id"`
	WeekNo      int    `json:"WeekNo"`
	HouseHoldId string `json:"HouseHoldId"`
	Meals       []Meal `json:"Meals"`
}

type PlanId struct {
	PlanId      int64  `json:"PlanId"`
	HouseholdId string `json:"HouseholdId"`
}

type MealId struct {
	MealId      int64  `json:"MealId"`
	HouseholdId string `json:"HouseholdId"`
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

type GeneratePlan struct {
	Id          int64   `json:"id"`
	WeekNo      int     `json:"WeekNo"`
	HouseholdId string  `json:"HouseholdId"`
	MealAmount  int     `json:"MealAmount"`
	Meals       []int64 `json:"Meals"`
}
