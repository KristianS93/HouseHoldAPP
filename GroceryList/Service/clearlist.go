package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"grocerylist/service/assistants"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// ClearList take a ListId, in a json body from a DELETE request, based on this list id, it deletes all items associated with the ListId
func (s *Server) ClearList(w http.ResponseWriter, r *http.Request) {

	//In any case return a json format
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	// get list id
	var rli assistants.RecieveId
	err := json.NewDecoder(r.Body).Decode(&rli)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error recieving list"}`)
		return
	}

	//no list was provided
	if rli.ListId == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No list provided"}`)
		return
	}

	//Instantiate a connection to mongo to list collection
	var itemClient database.MongClient
	itemClient.DbConnect(database.ConstGroceryItemsCollection)

	//Create a filter to the find query
	filter := bson.D{{Key: "ListId", Value: rli.ListId}}
	var results bson.D

	//findone query
	_ = itemClient.Connection.FindOne(context.TODO(), filter).Decode(&results)
	if results == nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "List is empty"}`)
		return
	}

	//Deletemany query
	_, err = itemClient.Connection.DeleteMany(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error deleting items in list"}`)
		return
	}
	defer itemClient.DbDisconnect()

	//Create succesfull json response
	str := make(map[string]string)
	str["Succes"] = "List Cleared"
	json.NewEncoder(w).Encode(str)

}
