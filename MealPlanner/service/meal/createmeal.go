package meal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

// CreateMeal takes POST request with a json body, determining the meal name, the week associated with it and the items in the meal.
func CreateMeal(w *http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(w)
	assistants.SetHeader(w)

	//If the method is not post, return wrong method
	//take the request pointer, pointer to response writer and the desired method.
	if assistants.WrongMethod(r, w, http.MethodPost) {
		log.Println("Wrong method, return 405")
		(*w).WriteHeader(405)
		io.WriteString((*w), `{"Error": "Bad method: wrong method"}`)
		return
	}

	//Get the json body and populate meal struct
	var mealformat models.Meal
	err := json.NewDecoder(r.Body).Decode(&mealformat)
	if err != nil {
		log.Println("Json didnt parse correct")
		(*w).WriteHeader(400)
		io.WriteString((*w), `{"Error": "Bad request: Getting data"}`)
		return
	}

	//Check that mealname and weekplanid is not missing
	if mealformat.MealName == "" {
		log.Println("Missing mealname")
		(*w).WriteHeader(400)
		io.WriteString((*w), `{"Error": "Bad request: Missing data"}`)
		return
	}

	err = db.InsertMeal(&mealformat)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Du har tilføjet et meal.")
}
