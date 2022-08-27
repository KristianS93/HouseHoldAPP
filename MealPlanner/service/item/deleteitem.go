package item

import (
	"encoding/json"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Decode json
	var getIds models.ItemIds
	err := json.NewDecoder(r.Body).Decode(&getIds)
	if err != nil {
		log.Println("Error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Decoding json"}`)
		return
	}

	//Checking if any there is provided any item ids
	if len(getIds.ItemIds) == 0 {
		log.Println("No ids provided")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "No ids provided"}`)
		return
	}

	//Delete items
	err = db.DeleteItems(getIds.ItemIds)
	if err != nil {
		log.Println("Error deleting items,", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Deleting items, items might not exist/already deleted"}`)
		return
	}

	//Return succes message
	log.Println("Item/items deleted")
	str := make(map[string]string)
	str["Succes"] = "Item/Items deleted"
	json.NewEncoder(w).Encode(str)

}
