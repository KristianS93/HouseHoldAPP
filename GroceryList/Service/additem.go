package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ListId   int    `json:"ListId"`
	Item     string `json:"ItemName"`
	Quantity int    `json:"Quantity"`
	Unit     string `json:"Unit"`
}

func (s *Server) AddItem(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	var xd []Item
	err := json.NewDecoder(r.Body).Decode(&xd)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Println(xd)
	fmt.Println(len(xd))
	// if r.Method != http.MethodPost {
	// 	w.WriteHeader(405)
	// 	return
	// }

	// //get list id
	// entries := strings.Split(r.URL.Path, "/")
	// if entries[2] != "id" {
	// 	//bad response
	// 	w.WriteHeader(400)
	// 	return
	// }

	// id, err := strconv.Atoi(entries[3])
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	return
	// }

	// //id now confirmed
	// fmt.Println(id)

} // /AddItem/id/2  || /AddItem?id=2
