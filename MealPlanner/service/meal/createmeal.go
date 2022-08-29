package meal

import (
	"encoding/json"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"mealplanner/service/item"
	"net/http"
)

// CreateMeal takes POST request with a json body, detmining the meal name, description and an array of items, it return a json reponse of the added meal id.
func CreateMeal(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Get the json body and populate meal struct
	var mealformat models.Meal
	err := assistants.DecodeData(r, &mealformat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	//Check that mealname and weekplanid is not missing
	if mealformat.MealName == "" {
		log.Println("Missing mealname")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Missing data"}`)
		return
	}

	var ids []int
	//Check that each item on the list also has a name
	if len(mealformat.Items) > 0 {
		for _, v := range mealformat.Items {
			if v.ItemName == "" {
				log.Println("Item missing a name")
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, `{"Error": "Bad request: Missing data"}`)
				return
			}
		}
		valueStr, valueArgs := item.FormatingInsertItems(mealformat.Items)

		//Create the items:
		ids, err = db.InsertItems(valueStr, valueArgs)
		if err != nil {
			log.Println("Error creating items")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"Error": "Bad request: Missing data"}`)
			return

		}
	}

	//Running insertmeal query
	err = db.InsertMeal(&mealformat, ids)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Creating meal"}`)
		return
	}

	//Creating json response
	type returndata struct {
		MealId int64 `json:"MealId"`
	}
	rd := returndata{mealformat.Id}
	log.Printf("Meal added")
	json.NewEncoder(w).Encode(rd)
}
