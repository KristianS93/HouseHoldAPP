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
	"strconv"
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

	i := 1
	var valuecon []string
	for j := 0; j < 2; j++ {
		var str []string
		temp := "("
		for k := 0; k < 3; k++ {
			str = append(str, "$"+strconv.Itoa(i))
			i++
		}
		temp += strings.Join(str, ",")
		temp += ")"
		valuecon = append(valuecon, temp)
	}
	valueStr := strings.Join(valuecon, ",")

	valueArgs := []interface{}{}
	// fmt.Println(count)
	for _, item := range getitems {
		valueArgs = append(valueArgs, item.ItemName)
		valueArgs = append(valueArgs, item.Quantity)
		valueArgs = append(valueArgs, item.Unit)
	}

	fmt.Println(valueArgs)
	fmt.Println(valueStr)

	err = db.InsertItems(valueStr, valueArgs)
	if err != nil {
		log.Println("Some error")
	}

}

//has to be able to insert multiple
// INSERT INTO household (householdid, meals, grocerylist) VALUES ('testhouse', '{1, 2, 4}', 'testlist')
// https://www.opsdash.com/blog/postgres-arrays-golang.html
