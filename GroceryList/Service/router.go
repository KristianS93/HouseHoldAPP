package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Routes handles the routing/handling of endpoints for the service
// The function is a method for the Server struct, and takes a pointer to a http.servemux type
func (s *Server) Routes(r *http.ServeMux) {
	r.HandleFunc("/GetList", s.GetList)
	r.HandleFunc("/AddItem/", s.AddItem)
	// r.HandleFunc("/DeleteItem/", s.DeleteItem)
	// r.HandleFunc("/ChangeItem/", s.ChangeItem)
}

func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {
	// Handle cors
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	//Check for correct method
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		io.WriteString(w, "Denied nigger")
		return
	}
	//correct logic to get the specified list

	//Serve the json for the list: json object with an array of items
	type item struct {
		Name   string
		Volume int
		Unit   string
	}

	type testdata []item
	var test2 testdata
	i1 := item{"Pasta", 1, "kg"}
	i2 := item{"Toilet Papir", 69, "pakke/pakker"}
	i3 := item{"Cancer treatment", 4, "RUNDER MED KEMO"}
	test2 = []item{i1, i2, i3}

	err := json.NewEncoder(w).Encode(test2)
	if err != nil {
		log.Println("Failed encoding data to JSON, error code: ", err)
	}
}
