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

type GetListId struct {
	ListId string `json:"ListId"`
}

func (s *Server) DeleteList(w http.ResponseWriter, r *http.Request) {
	//In any case return a json format
	assistants.EnableCors(&w)
	assistants.SetHeader(&w)

	//If the method is not delete, return wrong method
	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		io.WriteString(w, `{"Error": "Wrong method"}`)
		return
	}

	//Get the json body of the post and populate the dataforlist structure
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

	var client database.MongClient
	client.DbConnect(database.ConstGroceryListCollection)

	lookfor := dl.ListId

	filter := bson.D{{Key: "_id", Value: lookfor}}
	var existing bson.D

	//Checking if there is any matches on the list id, if not return 400
	_ = client.Connection.FindOne(context.TODO(), filter).Decode(&existing)
	if existing == nil {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "List does not exist"}`)
		return
	}

	_, err = client.Connection.DeleteOne(context.TODO(), filter)
	if err != nil {
		io.WriteString(w, `{"Error": "Failed deleting list"}`)
		w.WriteHeader(400)
		return
	}

	str := make(map[string]string)
	str["Succes"] = "List Deleted"
	json.NewEncoder(w).Encode(str)

}
