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

func CreateHouseHold(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	var householdid models.HouseHold
	err := assistants.DecodeData(r, &householdid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	//Check that mealname and weekplanid is not missing
	if householdid.HouseholdId == "" {
		log.Println("Missing household id")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Missing data"}`)
		return
	}

	householdid.GroceryListId = "No list"

	//Insert into household
	err = db.InsertHousehold(householdid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Internal Server error: inserting household"}`)
		return
	}

	//Create succesfull json response
	log.Println("Household added!")
	str := make(map[string]string)
	str["Succes"] = "Household created"
	json.NewEncoder(w).Encode(str)
}
