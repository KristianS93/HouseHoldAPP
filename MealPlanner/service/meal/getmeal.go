package meal

import (
	"encoding/json"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/service/assistants"
	"net/http"
	"strconv"
	"strings"
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
	mealid := r.URL.Query()["MealId"][0]

	//Check mealid is not empty
	if mealid == "" {
		log.Println("Url param not provided")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "no meal id provided"}`)
		return
	}

	//Check if meal id exist
	intMealId, err := strconv.Atoi(mealid)
	if err != nil {
		log.Println("Error converting mealid to int.")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Internal server error"}`)
		return
	}
	meal, itemString, err := db.SelectMealId(intMealId)
	if err != nil {
		log.Println("No meals found:", err)
	}

	itemIds := strings.Split(strings.TrimSuffix(itemString, ","), ",")

	//Get items
	items, err := db.SelectMultipleItems(itemIds)
	if err != nil {
		log.Println("Error selecting items")
		log.Println(err)
		log.Printf("%T", err)
	}
	meal.Items = items
	json.NewEncoder(w).Encode(meal)
}
