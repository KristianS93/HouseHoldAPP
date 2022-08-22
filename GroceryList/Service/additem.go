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

// CreateItem struct that can be used to create an item as a DTO
type CreateItem struct {
	ID       string `bson:"_id, omitempty"`
	ListId   string `bson:"ListId, omitempty" json:"ListId"`
	ItemName string `bson:"ItemName, omitempty" json:"ItemName"`
	Quantity string `bson:"Quantity" json:"Quantity"`
	Unit     string `bson:"Unit" json:"Unit"`
}

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
	var itemformat []CreateItem
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
	var itemInsertFormat []CreateItem
	for _, v := range itemformat {
		newId := primitive.NewObjectID()
		insertObj := CreateItem{string(newId.Hex()), v.ListId, v.ItemName, v.Quantity, v.Unit}
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
	
	//Create json response for succes.
	str := make(map[string]string)
	str["Succes"] = "Item Created"
	json.NewEncoder(w).Encode(str)

}
