package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// login is utilized internally to decode the request body
// from login attempts and other similar actions
type login struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type user struct {
	FirstName   string `json:"FirstName"`
	ListID      string `json:"ListID"`
	HouseholdID string `json:"HouseholdID"`
}

type NewUser struct {
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	ListID      string `json:"ListID"`
	HouseholdID string `json:"HouseholdID"`
}

func (s *Service) Routes(r *mux.Router) {
	r.HandleFunc("/Login", s.Login)
	r.HandleFunc("/CreateUser", s.CreateUser)

	// Create a User + Delete a User
	// Authorize Login Attempt
	//
}

// Login checks the LOGIN table for matching signatures
// of provided information and if appropriate login is provided,
// the user's name, listID and householdID is returned
// to the responsewriter, for the central server to handle
// and utilize in further requests and or front end display.
func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	var data login
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login: Failed to decode request body: ", err)
		return
	}

	if data.Email == "" || data.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload user
	err = s.Statements["Login"].QueryRow(data.Email, data.Password).Scan(&payload.FirstName, &payload.ListID, &payload.HouseholdID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			log.Println("Login: Record not found.")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login: Failed to access database: ", err)
		return
	}

	// this point is only reached when a record is found and no errors occured
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Println("Login: Failed to encode response.")
	}
}

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	var nu NewUser
	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateUser: Failed to decode request body: ", err)
		return
	}

	// must check database for already existing user before creating a duplicate, possibly refactor Login
	// should not check json, as stuff is internal - central server should check non empty etc. beforehand

	_, err = s.Statements["CreateUser"].Exec(nu.Email, nu.Password, nu.FirstName, nu.LastName, nu.ListID, nu.HouseholdID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateUser: Failed to access database and create new user, err: ", err)
		return
	}

	// this point is only reached when a new user is successfully registered, with no errors
	w.WriteHeader(http.StatusOK)
}
