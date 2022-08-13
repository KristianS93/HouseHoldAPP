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

func (s *Server) AddItem(w http.ResponseWriter, r *http.Request) {
	// #########################################
	//NOTE FIRST VERSION INSERTS 1 ITEM
	// #########################################

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
	var itemformat Item
	err := json.NewDecoder(r.Body).Decode(&itemformat)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	//Making sure data is not missing
	if itemformat.ListId == "" || itemformat.ItemName == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Missing data"}`)
		return
	}

	//Instantiate a connection to mongo
	var client database.MongClient
	client.DbConnect(database.ConstGroceryItemsCollection)

	//Consider Checking if an item is already
	//How ever this might not be important.

	//Generating item
	newId := primitive.NewObjectID()

	cItem := CreateItem{string(newId.Hex()),
		itemformat.ListId, itemformat.ItemName,
		itemformat.Quantity, itemformat.Unit}

	_, err = client.Connection.InsertOne(context.TODO(), cItem)
	if err != nil {
		io.WriteString(w, `{"Error": "Failed creating item"}`)
		w.WriteHeader(400)
		return
	}

	str := make(map[string]string)
	str["Succes"] = "Item Created"
	json.NewEncoder(w).Encode(str)

}
