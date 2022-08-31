package plan

import (
	"encoding/json"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
	"strconv"
)

func GetPlan(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Receive url query
	if len(r.URL.Query()) == 0 {
		log.Println("No url param exist")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No url exist}`)
		return
	}

	//Check if a url param called MealId exist
	if r.URL.Query()["weekno"] == nil {
		log.Println("Wrong url param")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Wrong url param"}`)
		return
	}
	if r.URL.Query()["householdid"] == nil {
		log.Println("Wrong url param")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Wrong url param"}`)
		return
	}
	householdid := r.URL.Query()["householdid"][0]
	//Convert weekno to int64
	weekno, err := strconv.ParseInt(r.URL.Query()["weekno"][0], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Could not convert weekno to int"}`)
		return
	}

	//Check that household and planid is not missing
	if weekno == 0 || householdid == "" {
		log.Println("No valid plan id or household id")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No valid plan or household id"}`)
		return
	}

	//Select data
	planData, err := db.SelectPlan(weekno, householdid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting plan"}`)
		return
	}

	//create a models.plan with json meals
	meals, err := db.SelectMultipleMeals(planData.Meals)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting meals"}`)
		return
	}

	var returnData models.ReturnPlan
	returnData.Id = planData.Id
	returnData.WeekNo = planData.WeekNo
	returnData.HouseHoldId = planData.HouseHoldId
	returnData.Meals = meals

	//Returning plan
	log.Printf("Plan returned")
	json.NewEncoder(w).Encode(returnData)
}
