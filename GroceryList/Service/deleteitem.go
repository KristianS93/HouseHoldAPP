package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type DeleteItem struct {
	ItemId string `json:"ItemId"`
}

func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// #########################################
	//NOTE FIRST VERSION DELETES 1 ITEM
	// #########################################

	//In any case return a json format
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	//If the method is not delete, return bad requst
	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Get the json body of the post and populate the Item structure
	var di DeleteItem
	err := json.NewDecoder(r.Body).Decode(&di)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Bad request: Getting data"}`)
		return
	}

	// //Make sure items id is not missing
	if di.ItemId == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Missing data"}`)
		return
	}

	//Instantiate a connection to mongo
	var client database.MongClient
	client.DbConnect(database.ConstGroceryItemsCollection)

	//Check if the if the items db has this item id
	lookfor := di.ItemId

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

	_, err = client.Connection.DeleteOne(context.TODO(), filter)
	if err != nil {
		io.WriteString(w, `{"Error": "Failed deleting item"}`)
		w.WriteHeader(400)
		return
	}

	str := make(map[string]string)
	str["Succes"] = "Item Deleted"
	json.NewEncoder(w).Encode(str)
}
