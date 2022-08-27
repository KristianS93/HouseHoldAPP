package item

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
	"strings"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Decode json
	var getitems []models.Item
	err := json.NewDecoder(r.Body).Decode(&getitems)
	if err != nil {
		log.Println("Error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Decoding json"}`)
		return
	}

	//Check itemname is not empty
	for _, v := range getitems {
		if v.ItemName == "" {
			log.Println("Missing item name")
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"Error": "Missing itemname"}`)
		}
	}

	//Generating value string in the format ($1, $2, $3),...
	i := 1
	var valuecon []string
	for j := 0; j < len(getitems); j++ {
		s := fmt.Sprintf("($%d, $%d, $%d)", i, i+1, i+2)
		valuecon = append(valuecon, s)
		i += 3
	}
	valueStr := strings.Join(valuecon, ",")

	//Generating a slice of interfaces with the values from each object.
	valueArgs := []interface{}{}
	for _, item := range getitems {
		valueArgs = append(valueArgs, item.ItemName)
		valueArgs = append(valueArgs, item.Quantity)
		valueArgs = append(valueArgs, item.Unit)
	}

	//Inserting items and returning a slice of itemids
	ids, err := db.InsertItems(valueStr, valueArgs)
	if err != nil {
		log.Println("Some error")
	}

	//Generating return data with an array of ids.
	type returndata struct {
		ItemIds []int `json:"ItemId"`
	}
	returnIds := returndata{}
	returnIds.ItemIds = append(returnIds.ItemIds, ids...)
	json.NewEncoder(w).Encode(returnIds)
}

// MAKING THE VALUE STR GENERIC!
// for j := 0; j < 2; j++ {
// 	var str []string
// 	temp := "("
// 	for k := 0; k < 3; k++ {
// 		str = append(str, "$"+strconv.Itoa(i))
// 		i++
// 	}
// 	temp += strings.Join(str, ",")
// 	temp += ")"
// 	valuecon = append(valuecon, temp)
// }
// valueStr := strings.Join(valuecon, ",")
// -----------------------------------------------

//has to be able to insert multiple
// INSERT INTO household (householdid, meals, grocerylist) VALUES ('testhouse', '{1, 2, 4}', 'testlist')
// https://www.opsdash.com/blog/postgres-arrays-golang.html
