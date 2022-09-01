package plan

import (
	"bytes"
	"encoding/json"
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
	if planData.WeekNo == 0 || planData.HouseHoldId == "" {
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
	//Hent specifik plan på hvert enkelt meal
	plan, err := db.SelectPlan(int64(planData.WeekNo), planData.HouseHoldId)
	if err != nil {
		log.Println("error selecting plan")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting household"}`)
		return
	}

	//get meal ids
	meals, err := db.SelectMultipleMeals(plan.Meals)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Selecting meals"}`)
		return
	}

	var itemIds []int64
	for _, v := range meals {
		itemIds = append(itemIds, v.Items...)
	}

	items, err := db.SelectMultipleItems(itemIds)
	if err != nil {
		log.Println("error selecting items")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "error selecting items"}`)
		return
	}

	var addItems []models.ItemRequest
	for _, v := range items {
		var item models.ItemRequest
		item.ListId = household.GroceryListId
		item.ItemName = v.ItemName
		item.Quantity = v.Quantity
		item.Unit = v.Unit
		addItems = append(addItems, item)
	}

	itemJson, err := json.Marshal(addItems)
	if err != nil {
		log.Println("error marshal item")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "error marshal items"}`)
		return

	}
	res, err := http.Post("http://localhost:5003/AddItem", "application/json", bytes.NewBuffer(itemJson))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "error posting to grocerylist"}`)
		return
	}

	if res.StatusCode != 200 {
		log.Println("error inserting into grocerylist")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "inserting into grocerylist"}`)
		return
	}
	//Succes
	//Return succes message
	log.Println("Grocery list Updated")
	str := make(map[string]string)
	str["Succes"] = "Grocery list Updated"
	json.NewEncoder(w).Encode(str)

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
