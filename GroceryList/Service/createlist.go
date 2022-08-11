package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type dataForList struct {
	WeekNo      int    `json:"WeekNo"`
	HouseholdId string `json:"HouseholdId"`
}

type createList struct {
	ListId      string
	WeekNo      int
	HouseholdId string
	Items       []Item
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
	fmt.Printf("HouseHold: %s, Weekno: %d, ListId: %s", cList.HouseholdId, cList.WeekNo, cList.ListId)
}
