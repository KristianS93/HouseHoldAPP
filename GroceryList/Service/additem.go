package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"grocerylist/service/assistants"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddItem Create item/items for a list, the function needs a post request, with a json object of an array of item/items
func (s *Server) AddItem(w http.ResponseWriter, r *http.Request) {
	//In any case return a json format
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//If the method is not post, return bad requst
	//take the request pointer, pointer to response writer and the desired method.
	if assistants.WrongMethod(r, &w, http.MethodPost) {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Bad method: wrong method"}`)
		return
	}

	//Get the json body of the post and populate the Item structure
	var itemformat []assistants.CreateItem
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

	//We insert the data obtained, into createitem datastructure, with a bson created id, for mongo.
	var itemInsertFormat []assistants.CreateItem
	for _, v := range itemformat {
		newId := primitive.NewObjectID()
		insertObj := assistants.CreateItem{ID: string(newId.Hex()), ListId: v.ListId, ItemName: v.ItemName, Quantity: v.Quantity, Unit: v.Unit}
		itemInsertFormat = append(itemInsertFormat, insertObj)
	}

	//To insert many each item has to be appended to slice of interface
	var insertItemQuery []interface{}
	for _, v := range itemInsertFormat {
		insertItemQuery = append(insertItemQuery, v)
	}

	//Insertmany query
	_, err = client.Connection.InsertMany(context.TODO(), insertItemQuery)
	if err != nil {
		io.WriteString(w, `{"Error": "Failed creating item"}`)
		w.WriteHeader(400)
		return
	}
	defer client.DbDisconnect()

	//Create json response for succes.
	str := make(map[string]string)
	str["Succes"] = "Item Created"
	json.NewEncoder(w).Encode(str)

}
