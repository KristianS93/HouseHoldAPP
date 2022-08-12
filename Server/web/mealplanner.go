package web

const (
	MealPlannerURL string = "localhost:5005"
)

type Item struct {
	Name     string
	Quantity string
	Unit     string
}

type MealPlan struct {
	Meals []Meal
}

type Meal struct {
	Name string
	// make a function that concats them quantity+unit+name
	Items   []Item
	Picture string
}

// func getMealPlan(userID string) MealPlan {
// 	// make an http request with the userID as the key
// 	// r, err := http.NewRequest(http.MethodGet, MealPlannerURL, nil)
// 	// if err != nil {
// 	// 	log.Println("getMealPlan-err:", err)
// 	// }

// }
