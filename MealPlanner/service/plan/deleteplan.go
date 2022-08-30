package plan

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

func DeletePlan(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Decode planid
	var getPlanId models.PlanId
	err := assistants.DecodeData(r, &getPlanId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	//Check that plan id is not empty
	if getPlanId.PlanId == 0 || getPlanId.HouseholdId == "" {
		log.Println("Plan or household id missing")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Missing household id or week number"}`)
		return
	}

	//Update household plans
	//Get household
	household, err := db.SelectHousehold(getPlanId.HouseholdId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting household"}`)
		return
	}

	//Remove the plan id
	planIds, err := assistants.RemoveIndex(household.Plans, getPlanId.PlanId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Plan doesnt exist in household"}`)
		return

	}
	fmt.Println(planIds)
	//Update meal
	err = db.UpdateHouseholdArrays(getPlanId.HouseholdId, "plans", planIds)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Updating household"}`)
		return

	}
	//Delete plan from plan db
	err = db.DeletePlan(getPlanId.PlanId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Deleting plan"}`)
		return
	}
	//Return Succes

	log.Println("Plan deleted")
	str := make(map[string]string)
	str["Succes"] = "Plan deleted"
	json.NewEncoder(w).Encode(str)
}
