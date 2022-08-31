package plan

import (
	"encoding/json"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

func ChangePlan(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Get the json body and populate meal struct
	var getPlan models.PlanDB
	err := assistants.DecodeData(r, &getPlan)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	//Check that household and week no is not empty
	if getPlan.HouseHoldId == "" || getPlan.WeekNo == 0 {
		log.Println("Household id or week number missing")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Missing household id or week number"}`)
		return
	}

	err = db.UpdatePlan(getPlan)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Updating plan"}`)
		return
	}

	//Return succes message
	log.Println("Plan Updated")
	str := make(map[string]string)
	str["Succes"] = "Plan Updated"
	json.NewEncoder(w).Encode(str)
}
