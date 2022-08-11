package service

import (
	"encoding/json"
	"fmt"
	"grocerylist/database"
	"net/http"
)

type deleteDate struct {
	ListId   int    `json:"ListId"`
	ItemName string `json:"ItemName"`
}

func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	var dd deleteDate
	err := json.NewDecoder(r.Body).Decode(&dd)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var client database.MongClient
	client.DbConnect()

	fmt.Printf("Item: %s is deleted from list id %d", dd.ItemName, dd.ListId)
}
