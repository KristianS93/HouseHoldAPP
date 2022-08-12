package service

import (
	"encoding/json"
	"io"
	"net/http"
)

type GetListId struct {
	ListId string `json:"ListId"`
}

func (s *Server) DeleteList(w http.ResponseWriter, r *http.Request) {
	//In any case return a json format
	w.Header().Set("Content-Type", "application/json")

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
		str := `{"Error": "Bad request"}`
		w.WriteHeader(400)
		io.WriteString(w, str)
		return
	}

	if dl.ListId == "" {
		w.WriteHeader(400)
		io.WriteString(w, `{"Error": "No list provided"}`)
		return
	}

	//Instantiate a connection to mongo
	// var client database.MongClient
	// client.DbConnect()

}
