package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type RecieveListId struct {
	ListId string `json:"ListId"`
}

func (s *Server) ClearList(w http.ResponseWriter, r *http.Request) {

	//In any case return a json format
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	// get list id
	var rli RecieveListId
	err := json.NewDecoder(r.Body).Decode(&rli)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error recieving list"}`)
		return
	}

	if rli.ListId == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No list provided"}`)
		return
	}

	//Instantiate a connection to mongo to list collection
	var itemClient database.MongClient
	itemClient.DbConnect(database.ConstGroceryItemsCollection)

	filter := bson.D{{Key: "ListId", Value: rli.ListId}}

	cur, err := itemClient.Connection.Find(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error finding list id in items"}`)
		return
	}

	var results []bson.D
	if err = cur.All(context.TODO(), &results); err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error writing items to array"}`)
		return
	}

	if len(results) == 0 {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "List is empty"}`)
		return
	}

	_, err = itemClient.Connection.DeleteMany(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "Error deleting items in list"}`)
		return
	}

	str := make(map[string]string)
	str["Succes"] = "List Cleared"
	json.NewEncoder(w).Encode(str)

}
