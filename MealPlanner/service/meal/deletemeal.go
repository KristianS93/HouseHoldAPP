package meal

import (
	"mealplanner/database"
	"mealplanner/service/assistants"
	"net/http"
)

func DeleteMeal(w *http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(w)
	assistants.SetHeader(w)

	//Decode MealId

}
