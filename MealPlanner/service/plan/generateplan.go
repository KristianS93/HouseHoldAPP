package plan

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

func GeneratePlan(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	var newPlan models.GeneratePlan
	err := assistants.DecodeData(r, &newPlan)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Error receiving data"}`)
		return
	}

	//check that week no and household id is not empty
	if newPlan.WeekNo == 0 || newPlan.HouseholdId == "" || newPlan.MealAmount == 0 {
		log.Println("No week no. or household id provided")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No week no. or household provided"}`)
		return
	}

	// get array of meals to the household
	household, err := db.SelectHousehold(newPlan.HouseholdId)
	if err != nil {
		log.Println("Error selecting household")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting household"}`)
		return
	}

	//Check that meal amount is lesser than amount of meals
	if newPlan.MealAmount > len(household.Meals) {
		log.Println("Error not enough meals to create plan")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Not enough meals to create plan"}`)
		return
	}

	//Generating array with random meals
	mealsList := randomMeals(newPlan.MealAmount, household.Meals)
	//Creating valid model for inserting into plan db
	var insertPlan models.Plan
	insertPlan.WeekNo = newPlan.WeekNo
	insertPlan.HouseHoldId = newPlan.HouseholdId

	//If weekno already is used return error
	if !db.TestMultipleWeekno(insertPlan.WeekNo, insertPlan.HouseHoldId) {
		log.Println("WeekNo already exist for this household")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "WeekNo already exist for this household"}`)
		return
	}
	//DETTE OVERHER SKAL ÆNDRES TIL ET NAVN I STEDET FOR, DET HER ER RETARDET

	// Insert into plan
	err = db.CreatePlan(&insertPlan, mealsList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Creating plan"}`)
		return
	}

	//Select the previous added plan ids
	getHousehold, err := db.SelectHousehold(newPlan.HouseholdId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting household"}`)
		return
	}
	getHousehold.Plans = append(getHousehold.Plans, insertPlan.Id)

	//Insert(UPDATE) the plan into the household to keep track of the plan
	err = db.UpdateHouseholdArrays(insertPlan.HouseHoldId, "plans", getHousehold.Plans)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Updating household"}`)
		return
	}

	//Return data
	type returndata struct {
		WeekNo int64 `json:"WeekNo"`
	}
	rd := returndata{int64(insertPlan.WeekNo)}
	log.Printf("plan added")
	json.NewEncoder(w).Encode(rd)
}

func randomMeals(mealAmount int, mealList []int64) []int64 {
	var newMap = make(map[int64]int64)
	for len(newMap) < mealAmount {
		randomN := rand.Intn(len(mealList)-0) + 0
		if newMap[int64(randomN)] == 0 {
			newMap[int64(randomN)] = int64(randomN)
		}
	}
	var returnList []int64
	for _, v := range newMap {
		returnList = append(returnList, v)
	}
	return returnList
}
