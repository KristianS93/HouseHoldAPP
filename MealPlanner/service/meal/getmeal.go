package meal

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

func GetMeal(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Receive url query
	if len(r.URL.Query()) == 0 {
		log.Println("No url param exist")
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, `{"Error": "No url exist}`)
		return
	}

	//Check if a url param called MealId exist
	if r.URL.Query()["MealId"] == nil {
		log.Println("Wrong url param")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Wrong url param"}`)
		return
	}

	//Get mealid
	mealid, err := strconv.ParseInt(r.URL.Query()["MealId"][0], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Could not convert meal id to int"}`)
		return
	}

	//Check mealid is not empty
	if mealid == 0 {
		log.Println("Url param not provided")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "no meal id provided"}`)
		return
	}

	meal, err := db.SelectMealId(mealid)
	if err != nil {
		log.Println("No meals found:", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No meals found"}`)
		return
	}

	//Get items
	items, err := db.SelectMultipleItems(meal.Items)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Error selecting items"}`)
		return
	}
	fmt.Println(items)
	var returnData models.Meal
	returnData.Id = meal.Id
	returnData.MealName = meal.MealName
	returnData.Description = meal.Description
	returnData.Items = items

	json.NewEncoder(w).Encode(returnData)
}
