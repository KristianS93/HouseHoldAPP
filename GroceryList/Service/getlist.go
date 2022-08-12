package service

import (
	"context"
	"encoding/json"
	"fmt"
	"grocerylist/database"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	//Get the json body of the post and populate the GetListId structure
	var dl GetListId
	err := json.NewDecoder(r.Body).Decode(&dl)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Bad request"}`)
		return
	}

	if dl.ListId == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No list provided"}`)
		return
	}

	//Instantiate a connection to mongo
	var client database.MongClient
	client.DbConnect(database.ConstGroceryItemsCollection)

	//Get data from list
	lookfor := dl.ListId
	filter := bson.D{primitive.E{Key: "ListId", Value: lookfor}}

	res, err := client.Connection.Find(context.TODO(), filter)
	if res != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No items on the list"}`)
		return
	}

	var itemsList []Item
	if err = res.All(context.TODO(), &itemsList); err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Could not retrieve list"}`)
	}

	fmt.Println("displaying all results from the search query")
	for _, result := range itemsList {
		fmt.Println(result)
	}

	// err := json.NewEncoder(w).Encode(test2)
	// if err != nil {
	// 	log.Println("Failed encoding data to JSON, error code: ", err)
	// }
}

//######################
// result.InsertedID
// lookfor := idString
// var results bson.M
// filter := bson.D{{"_id", lookfor}}
// if err = client.Connection.FindOne(context.TODO(), filter).Decode(&results); err != nil {
// 	panic(err)
// }
// fmt.Println(results)
//#################
