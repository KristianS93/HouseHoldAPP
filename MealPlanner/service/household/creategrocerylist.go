package household

import (
	"encoding/json"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

func CreateGroceryList(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	var household models.HouseHold
	err := assistants.DecodeData(r, &household)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}
	//Check household id is not empty
	if household.HouseholdId == "" || household.GroceryListId == "" {
		log.Println("No household/grocerylist provided")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: No household/grocerylist provided"}`)
		return
	}

	//Check and receive
	_, err = db.SelectGroceryList(household)
	if err != nil {
		log.Println("No household", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No Household in database"}`)
		return
	}

	//update grocerylist
	err = db.UpdateHousehold(household, "grocerylist", "")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Updating grocerylist id"}`)
		return
	}

	//Succes
	log.Println("Grocerylist added!")
	str := make(map[string]string)
	str["Succes"] = "Grocerylist added"
	json.NewEncoder(w).Encode(str)
}
