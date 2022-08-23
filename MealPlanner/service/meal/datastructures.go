package meal

type Item struct {
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

type Meal struct {
	MealName   string `json:"MealName"`
	WeekPlanId string `json:"WeekPlanId"`
	Items      []Item `json:"Items"`
}

type WeekPlan struct {
	WeekNo      string `json:"WeekNo"`
	HouseHoldId string `json:"HouseHoldId"`
	Meals       []Meal `json:"Meals"`
}
