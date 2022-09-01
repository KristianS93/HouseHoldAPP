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

func DeleteHouseHold(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Receive json
	var household models.HouseHold
	err := assistants.DecodeData(r, &household)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Failed decoding json body"}`)
		return
	}

	//check that household id is not empty
	if household.HouseholdId == "" {
		log.Println("No household id provided")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: No household id provided"}`)
		return
	}

	// //Check and receive
	getHousehold, err := db.SelectHousehold(household.HouseholdId)
	if err != nil {
		log.Println("No household", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No Household in database"}`)
		return
	}

	//Get all items meals
	meals, err := db.SelectMultipleMeals(getHousehold.Meals)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting meals"}`)
		return
	}

	// ALL item ids
	var itemIds []int64
	for _, v := range meals {
		itemIds = append(itemIds, v.Items...)
	}

	//Delete all items
	err = db.DeleteItems(itemIds)
	if err != nil {
		log.Println("error deleting items")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "deleting items"}`)
		return
	}

	//delete meals
	err = db.DeleteMeals(getHousehold.Meals)
	if err != nil {
		log.Println("error deleting meals")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "deleting meals"}`)
		return
	}

	//Delete plans
	err = db.DeletePlans(getHousehold.Plans)
	if err != nil {
		log.Println("error deleting plans")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "deleting plans"}`)
		return
	}

	err = db.DeleteHouseHold(getHousehold.HouseholdId)
	if err != nil {
		log.Println("error deleting household")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "deleting household"}`)
		return
	}

	//Return succes message
	log.Println("Household deleted")
	str := make(map[string]string)
	str["Succes"] = "Household deleted"
	json.NewEncoder(w).Encode(str)
}
