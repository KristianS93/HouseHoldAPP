package plan

import (
	"fmt"
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
	
	//Delete plan from plan db
	
	//Return Succes	
}
