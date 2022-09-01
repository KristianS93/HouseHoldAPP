package plan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

func GenerateList(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//get household and plan
	var planData models.PlanDB
	err := assistants.DecodeData(r, &planData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Error receiving data"}`)
		return
	}

	//Check specfic fields are not empty
	if planData.WeekNo == 0 || planData.HouseHoldId == "" || len(planData.Meals) == 0 {
		log.Println("No week no. or household id provided or meals")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No week no. or household provided or meals"}`)
		return
	}
	//Get list id from household
	household, err := db.SelectHousehold(planData.HouseHoldId)
	if err != nil {
		log.Println("Error selecting household")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting household"}`)
		return
	}

	//First we need to clear the list if it exist, BE SURE YOU AWARE THE USER OF THIS BEFORE
	errStr, err := clearList(household.GroceryListId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "`+errStr+`"}`)
		return
	}
	//Hent items på hvert enkelt meal
	fmt.Println("list cleared")
	//add items til list

	//Succes

	fmt.Println("Noget med en plan")
}

func clearList(listId string) (string, error) {
	type listid struct {
		ListId string `json:"ListId"`
	}
	li := listid{listId}
	liJson, err := json.Marshal(li)
	if err != nil {
		log.Println("error marshalling json")
	}
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:5003/ClearList", bytes.NewBuffer(liJson))
	if err != nil {
		return "error creating delete request", err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "failed delete request", err
	}

	if res.StatusCode != http.StatusOK {
		return "could not clear list", err
	}
	return "", nil
}
