package service

import (
	"context"
	"encoding/json"
	"grocerylist/database"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type dataForList struct {
	WeekNo      int    `json:"WeekNo"`
	HouseholdId string `json:"HouseholdId"`
}

type createList struct {
	ListId      string `bson:"ListId,omitempty"`
	WeekNo      int    `bson:"WeekNo"`
	HouseholdId string `bson:"HouseholdId, omitempty"`
	Items       []Item `bson:"Items"`
}

func (s *Server) CreateList(w http.ResponseWriter, r *http.Request) {

	var cl dataForList
	err := json.NewDecoder(r.Body).Decode(&cl)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var cList createList
	strListid := "list:" + uuid.New().String()
	cList.ListId = strListid
	cList.WeekNo = cl.WeekNo
	cList.HouseholdId = cl.HouseholdId

	var client database.MongClient
	client.DbConnect()

	_, err = client.Connection.InsertOne(context.TODO(), cList)
	if err != nil {
		log.Fatalln("Could not create list", err)
		w.WriteHeader(400)
		return
	}

	json.NewEncoder(w).Encode(cList.ListId)
}
