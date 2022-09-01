package household

import (
	"net/http"
)

//household, grocerylist list er lavet.
//

func CreateGroceryList(w http.ResponseWriter, r *http.Request) {
	// db := database.Connect()
	// defer db.Con.Close()
	// //Enable cors and set header to return json
	// assistants.EnableCors(&w)
	// assistants.SetHeader(&w)

	// var household models.HouseHold
	// err := assistants.DecodeData(r, &household)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
	// 	return
	// }

	// //Check household id is not empty
	// if household.HouseholdId == "" {
	// 	log.Println("No household")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "No household"}`)
	// 	return
	// }

	// //Check and receive
	// getHousehold, err := db.SelectHousehold(household.HouseholdId)
	// if err != nil {
	// 	log.Println("No household", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "No Household in database"}`)
	// 	return
	// }

	// //Check if a list is active or not
	// if getHousehold.GroceryListId != "No list" {
	// 	log.Println("grocery list already added")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "grocery list already added"}`)
	// 	return
	// }

	// //Call creategrocerylist on grocerylist service
	// reqGroceryList := models.HouseholdRequest{HouseholdId: getHousehold.HouseholdId}
	// jsonData, err := json.Marshal(reqGroceryList)
	// if err != nil {
	// 	log.Println("error marshaling data")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "Marshaling data"}`)
	// 	return
	// }
	// res, err := http.Post("http://localhost:5003/CreateList", "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	log.Println("error posting to grocerylist service")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "posting to grocerylist service"}`)
	// 	return
	// }
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Println("error reading body")
	// }
	// if res.StatusCode != http.StatusOK {
	// 	log.Println("Recieved error code")
	// 	log.Println(string(data))
	// }

	// var getListId models.HouseholdGroceryList
	// fmt.Println(string(data))

	// err = json.Unmarshal(data, &getListId)
	// if err != nil {
	// 	log.Println("error decoding list id")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "decoding list id"}`)
	// 	return
	// }

	// //update grocerylist
	// var householdUpdate models.HouseHold
	// householdUpdate.HouseholdId = getHousehold.HouseholdId
	// householdUpdate.GroceryListId = getListId.ListId
	// err = db.UpdateHouseholdGroceryList(householdUpdate)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "Updating grocerylist id"}`)
	// 	return
	// }

	// //Succes
	// log.Println("Grocerylist added!")
	// str := make(map[string]string)
	// str["Succes"] = "Grocerylist added"
	// json.NewEncoder(w).Encode(str)
}
