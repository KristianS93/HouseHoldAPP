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

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Decode json
	var getitems models.Item
	err := json.NewDecoder(r.Body).Decode(&getitems)
	if err != nil {
		log.Println("Error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Decoding json"}`)
		return
	}

	//Check itemname or id is not empty
	if getitems.ItemName == "" || getitems.Id == 0 {
		log.Println("Missing item name / no ide provided")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Missing itemname or id"}`)
		return
	}

	err = db.UpdateItem(getitems)
	if err != nil {
		log.Println("Error updating item")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Error": "Updating item"}`)
		return
	}

	//Return succes message
	log.Println("Item updated")
	str := make(map[string]string)
	str["Succes"] = "Item updated"
	json.NewEncoder(w).Encode(str)

}
