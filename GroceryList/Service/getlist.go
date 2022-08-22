package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"grocerylist/service/assistants"
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
	ID       string `bson:"_id"`
	ItemName string `json:"ItemName"`
	Quantity string `json:"Quantity"`
	Unit     string `json:"Unit"`
}

// GetList takes a URL parameter called ?ListId=xxx from a GET request, an returns a json object with all items associated to this list id from the mongo db.
func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {

	//In any case return a json format and enable cors
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//Check for correct method
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Check that a url parameter exist.
	if len(r.URL.Query()) == 0 {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "No list provided"}`)
		return
	}

	//Check if a url parameter called ListId exist
	if r.URL.Query()["ListId"] == nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Wrong url param"}`)
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

	//Find query based on the filter
	res, err := client.Connection.Find(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No items on the list"}`)
		return
	}
	defer client.DbDisconnect()

	//Insert data from db into ItemList datastructure, so it can be encoded into json
	var itemsList []ItemList
	if err = res.All(context.TODO(), &itemsList); err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Could not retrieve list"}`)
	}

	//Encode data from itemlist into json.
	var returndata = ItemHolder{"List Retrieved", itemsList}
	json.NewEncoder(w).Encode(returndata)
}
