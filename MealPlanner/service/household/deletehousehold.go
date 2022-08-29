package household

import (
	"net/http"
)

func DeleteHouseHold(w http.ResponseWriter, r *http.Request) {
	// db := database.Connect()
	// defer db.Con.Close()
	// //Enable cors and set header to return json
	// assistants.EnableCors(&w)
	// assistants.SetHeader(&w)

	// //Receive json
	// var household models.HouseHold
	// err := assistants.DecodeData(r, &household)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	io.WriteString(w, `{"Error": "Internal servererror: Failed decoding json body"}`)
	// 	return
	// }

	// //check that household id is not empty
	// if household.HouseholdId == "" {
	// 	log.Println("No household id provided")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "Bad request: No household id provided"}`)
	// 	return
	// }

	// //Check and receive
	// var getHousehold models.HouseHold
	// _, err = db.SelectHousehold(getHousehold)
	// if err != nil {
	// 	log.Println("No household", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, `{"Error": "No Household in database"}`)
	// 	return
	// }

	//if there is meals associated or a grocerylist, these has to be deleted.
	//Kald grocery list api clearlist og deletelist
	//Delete meals + items.
}
