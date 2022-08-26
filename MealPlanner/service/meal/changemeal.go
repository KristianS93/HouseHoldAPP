package meal

import (
	"log"
	"mealplanner/database"
	"mealplanner/service/assistants"
	"net/http"
)

func ChangeMeal(w *http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(w)
	assistants.SetHeader(w)

	log.Println("Du har tilføjet et meal.")
}
