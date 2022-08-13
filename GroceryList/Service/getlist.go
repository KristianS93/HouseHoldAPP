package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemHolder struct {
	Succes string     `json:"Succes"`
	Items  []ItemList `json:"Items"`
}
type ItemList struct {
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {

	//In any case return a json format and enable cors
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	//Check for correct method
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Receive list id from s
	recievedListId := r.URL.Query()["ListId"][0]

	//Check that list id is not empty
	if recievedListId == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No list provided"}`)
		return
	}

	//Instantiate a connection to mongo to list collection
	var listClient database.MongClient
	listClient.DbConnect(database.ConstGroceryListCollection)

	//check that list exist
	// lookfor := di.ItemName
	filterList := bson.D{{Key: "_id", Value: recievedListId}}
	var resultList bson.D

	//Checking if there is any matches on the house hold id, if so return 400
	_ = listClient.Connection.FindOne(context.TODO(), filterList).Decode(&resultList)
	if resultList == nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "List does not exist"}`)
		return
	}

	//Instantiate a connection to mongo items collection
	var client database.MongClient
	client.DbConnect(database.ConstGroceryItemsCollection)

	//Get data from list
	lookfor := recievedListId
	filter := bson.D{primitive.E{Key: "ListId", Value: lookfor}}

	res, err := client.Connection.Find(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No items on the list"}`)
		return
	}

	var itemsList []ItemList
	if err = res.All(context.TODO(), &itemsList); err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Could not retrieve list"}`)
	}

	var returndata = ItemHolder{"List Retrieved", itemsList}

	json.NewEncoder(w).Encode(returndata)
}

// //Get the json body of the post and populate the GetListId structure
// var dl GetListId
// err := json.NewDecoder(r.Body).Decode(&dl)
// if err != nil {
// 	w.WriteHeader(400)
// 	io.WriteString(w, `{"Error": "Bad request"}`)
// 	return
// }

// if dl.ListId == "" {
// 	w.WriteHeader(400)
// 	io.WriteString(w, `{"Error": "No list provided"}`)
// 	return
// }
