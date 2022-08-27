package household

import (
	"io"
	"mealplanner/database"
	"mealplanner/models"
	"mealplanner/service/assistants"
	"net/http"
)

func CreateGroceryList(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Con.Close()
	//Enable cors and set header to return json
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	var household models.HouseHold
	err := assistants.DecodeData(r, &household)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

}
