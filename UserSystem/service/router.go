package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type login struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (s *Service) Routes(r *mux.Router) {
	r.HandleFunc("/CheckLogin", s.CheckLogin)
}

func (s *Service) CheckLogin(w http.ResponseWriter, r *http.Request) {
	var user login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("failed to decode: ", err)
	}

	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	entry, err := s.Statements["CheckLogin"].Query(user.Email, user.Password)
	if err != nil {
		log.Println("Failed to access database: ", err)
	}
	defer entry.Close()

	// return true if record exists
	if entry.Next() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
