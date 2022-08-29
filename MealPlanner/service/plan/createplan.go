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

func CreatePlan(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Get the json body and populate meal struct
	var getPlan models.Plan
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

	//check that if meals slice is bigger then 0 we have to check this aswell
	var mealIds []int64
	if len(getPlan.Meals) > 0 {
		//Handle meals, ids are available at this point.
		for _, v := range getPlan.Meals {
			mealIds = append(mealIds, v.Id)
		}
	}

	err = db.CreatePlan(&getPlan, mealIds)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Creating plan"}`)
		return
	}

	//Select the previous added plan ids
	getHousehold, err := db.SelectHousehold(getPlan.HouseHoldId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting household"}`)
		return

	}
	getHousehold.Plans = append(getHousehold.Plans, getPlan.Id)
	//Insert(UPDATE) the plan into the household to keep track of the plan
	err = db.UpdateHouseholdArrays(getPlan.HouseHoldId, "plans", getHousehold.Plans)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Updating household"}`)
		return
	}

	//Creating json response //MAYBE RETURN HOUSEHOLD ID IF IT IS OF USE TO THE FRONTEND
	type returndata struct {
		PlanId int64 `json:"Planid"`
	}
	rd := returndata{getPlan.Id}
	log.Printf("Plan added")
	json.NewEncoder(w).Encode(rd)
}
