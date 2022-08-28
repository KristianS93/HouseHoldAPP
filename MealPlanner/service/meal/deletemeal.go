package meal

import (
	"encoding/json"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

func DeleteMeal(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Decode MealId
	var getMealId models.MealId
	err := json.NewDecoder(r.Body).Decode(&getMealId)
	if err != nil {
		log.Println("Error decoding mealid")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Decoding json"}`)
		return
	}

	//Check that a mealid is provided
	if getMealId.MealId == 0 {
		log.Println("No meal id provided")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "No meal id provided"}`)
		return
	}

	//Get ids of all items on the mealid
	getMeal, err := db.SelectMealId(getMealId.MealId)
	if err != nil {
		log.Println(err)
		log.Println("Error selecting meal")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Selecting meal"}`)
		return
	}

	//Delete items
	err = db.DeleteItems(getMeal.Items)
	if err != nil {
		log.Println("Error deleting items")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Deleting items in meal"}`)
		return
	}

	//Delete meal
	err = db.DeleteMeal(getMealId.MealId)
	if err != nil {
		log.Println("Error deleting meal")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Deleting meal"}`)
		return
	}

	//Return succes message
	log.Println("Meal deleted")
	str := make(map[string]string)
	str["Succes"] = "Meal deleted"
	json.NewEncoder(w).Encode(str)
}
