package service

import (
	"context"
	"encoding/json"
	"fmt"
	"grocerylist/database"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type getItem struct {
	ID       string `json:"Id"`
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}
type itemChange struct {
	ID       string `bson:"_id"`
	ItemName string `bson:"ItemName"`
	Quantity string `bson:"Quantity"`
	Unit     string `bson:"Unit"`
}

func (s *Server) ChangeItem(w http.ResponseWriter, r *http.Request) {
	//In any case return a json format
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	//If the method is not delete, return bad requst
	if r.Method != http.MethodPatch {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Get the json body of the post and populate the Item structure
	var ui getItem
	err := json.NewDecoder(r.Body).Decode(&ui)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	// //Make sure items id is not missing
	if ui.ID == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Missing data"}`)
		return
	}

	//Instantiate a connection to mongo
	var client database.MongClient
	client.DbConnect(database.ConstGroceryItemsCollection)

	//Making sure the item exist
	lookfor := ui.ID

	// lookfor := di.ItemName
	filter := bson.D{{Key: "_id", Value: lookfor}}
	var results bson.D

	//Checking if there is any matches on the house hold id, if so return 400
	_ = client.Connection.FindOne(context.TODO(), filter).Decode(&results)
	if results == nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Item does not exist"}`)
		return
	}

	//Item exist, update item

	changeItem := itemChange{ui.ID, ui.ItemName, ui.Quantity, ui.Unit}

	updateItem := bson.D{{Key: "$set", Value: changeItem}}

	_, err = client.Connection.UpdateOne(context.TODO(), filter, updateItem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error updating item"}`)
		return
	}

	str := make(map[string]string)
	str["Succes"] = "Item Updated"
	json.NewEncoder(w).Encode(str)

}
