package service

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) AddItem(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	//get list id
	entries := strings.Split(r.URL.Path, "/")
	if entries[2] != "id" {
		//bad response
		w.WriteHeader(400)
		return
	}

	id, err := strconv.Atoi(entries[3])
	if err != nil {
		w.WriteHeader(400)
		return
	}

	//id now confirmed
	fmt.Println(id)

}
