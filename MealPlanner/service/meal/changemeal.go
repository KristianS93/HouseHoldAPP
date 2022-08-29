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

func ChangeMeal(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Decode json
	var getMeal models.Meal
	err := assistants.DecodeData(r, &getMeal)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	if getMeal.Id == 0 || getMeal.MealName == "" {
		log.Println("Error, not id or mealname provided")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Missing id or mealname"}`)
		return
	}

	//Update Query
	err = db.UpdateMeal(getMeal)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Updating meal}`)
		return
	}

	//Return succes message
	log.Println("Meal Updated")
	str := make(map[string]string)
	str["Succes"] = "Meal Updated"
	json.NewEncoder(w).Encode(str)
}
