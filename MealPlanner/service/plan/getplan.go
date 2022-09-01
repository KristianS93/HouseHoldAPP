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

	// ALL item ids
	var itemIds []int64
	for _, v := range meals {
		itemIds = append(itemIds, v.Items...)
	}

	itemMap := make(map[int64]models.Item)

	items, err := db.SelectMultipleItems(itemIds)
	if err != nil {
		log.Println("error selecting items")
	}

	for _, v := range items {
		var item models.Item
		item.Id = v.Id
		item.ItemName = v.ItemName
		item.Quantity = v.Quantity
		item.Unit = v.Unit
		itemMap[int64(item.Id)] = item
	}

	var mealsReturn []models.Meal
	for _, v := range meals {
		var meal models.Meal
		meal.Id = v.Id
		meal.HouseholdId = householdid
		meal.MealName = v.MealName
		meal.Description = v.Description
		fmt.Println(len(v.Items))
		var newItems []models.Item
		for _, x := range v.Items {
			newItems = append(newItems, itemMap[x])
		}
		meal.Items = newItems
		mealsReturn = append(mealsReturn, meal)
	}

	// fmt.Println(mealsReturn)
	// meals.Items

	var returnData models.ReturnPlan
	returnData.Id = planData.Id
	returnData.WeekNo = planData.WeekNo
	returnData.HouseHoldId = planData.HouseHoldId
	returnData.Meals = mealsReturn

	//Returning plan
	log.Printf("Plan returned")
	json.NewEncoder(w).Encode(returnData)
}
