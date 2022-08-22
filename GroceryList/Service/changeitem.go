package service

import (
	"context"
	"encoding/json"
	"fmt"
	"grocerylist/database"
	"grocerylist/service/assistants"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// ChangeItem takes a json object in the ItemChange format, KAN ÆNDRES TIL CREATEITEM SOM MÅSKE KAN BLIVE HANDLEITEM, and updates the desired item, based on the item id.
func (s *Server) ChangeItem(w http.ResponseWriter, r *http.Request) {
	//In any case return a json format
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//If the method is not delete, return bad requst
	if r.Method != http.MethodPatch {
		fmt.Println("method")
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Get the json body of the post and populate the Item structure
	var ui assistants.CreateItem
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
	changeItem := assistants.ItemData{Id: ui.ID, ItemName: ui.ItemName, Quantity: ui.Quantity, Unit: ui.Unit}

	updateItem := bson.D{{Key: "$set", Value: changeItem}}

	//Update one query
	_, err = client.Connection.UpdateOne(context.TODO(), filter, updateItem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error updating item"}`)
		return
	}
	defer client.DbDisconnect()

	//Create json, item created.
	str := make(map[string]string)
	str["Succes"] = "Item Updated"
	json.NewEncoder(w).Encode(str)

}
