package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ListId   string `json:"ListId"`
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

type CreateItem struct {
	ID       string `bson:"_id, omitempty"`
	ListId   string `bson:"ListId, omitempty"`
	Item     string `bson:"ItemName, omitempty"`
	Quantity string `bson:"Quantity"`
	Unit     string `bson:"Unit"`
}

// AddItem Create item/items for a list, the function needs a post request, with a json object of an array of item/items
func (s *Server) AddItem(w http.ResponseWriter, r *http.Request) {
	//In any case return a json format
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	//If the method is not post, return bad requst
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Get the json body of the post and populate the Item structure
	var itemformat []Item
	err := json.NewDecoder(r.Body).Decode(&itemformat)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	//Making sure data is not missing from any of the items
	for _, v := range itemformat {
		// i = index v = value hvilket er item her
		if v.ListId == "" || v.ItemName == "" {
			w.WriteHeader(400)
			io.WriteString(w, `{"Error": "Missing data"}`)
			return
		}
	}

	//Instantiate a connection to mongo
	var client database.MongClient
	client.DbConnect(database.ConstGroceryItemsCollection)

	var itemInsertFormat []CreateItem

	for _, v := range itemformat {
		newId := primitive.NewObjectID()
		insertObj := CreateItem{string(newId.Hex()), v.ListId, v.ItemName, v.Quantity, v.Unit}
		itemInsertFormat = append(itemInsertFormat, insertObj)
	}

	var insertItemQuery []interface{}
	for _, v := range itemInsertFormat {
		insertItemQuery = append(insertItemQuery, v)
	}

	_, err = client.Connection.InsertMany(context.TODO(), insertItemQuery)
	if err != nil {
		io.WriteString(w, `{"Error": "Failed creating item"}`)
		w.WriteHeader(400)
		return
	}

	str := make(map[string]string)
	str["Succes"] = "Item Created"
	json.NewEncoder(w).Encode(str)

}
