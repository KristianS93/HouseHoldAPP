package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var UniqueError string = "UNIQUE constraint failed: USERS.userID"

// login is utilized internally to decode the request body
// from login attempts and other similar actions
type login struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
}

type user struct {
	FirstName   string `json:"FirstName"`
	ListID      string `json:"ListID"`
	HouseholdID string `json:"HouseholdID"`
}

type NewUser struct {
	UserID      string `json:"UserID"`
	Password    string `json:"Password"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	ListID      string `json:"ListID"`
	HouseholdID string `json:"HouseholdID"`
}

func (s *Service) Routes(r *mux.Router) {
	r.HandleFunc("/Login", s.Login)
	r.HandleFunc("/CreateUser", s.CreateUser)
	r.HandleFunc("/DeleteUser", s.DeleteUser)

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

	var payload user
	err = s.Statements["Login"].QueryRow(data.UserID, data.Password).Scan(&payload.FirstName, &payload.ListID, &payload.HouseholdID)
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

	_, err = s.Statements["CreateUser"].Exec(nu.UserID, nu.Password, nu.FirstName, nu.LastName, nu.ListID, nu.HouseholdID)
	if err != nil {
		if err.Error() == UniqueError {
			w.WriteHeader(http.StatusConflict)
			log.Println("DeleteUser: user already exists")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateUser: Failed to access database and create new user, err: ", err)
		return
	}

	// this point is only reached when a new user is successfully registered, with no errors
	w.WriteHeader(http.StatusCreated)
}

func (s *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var data login
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("DeleteUser: Failed to decode request body: ", err)
		return
	}

	result, err := s.Statements["DeleteUser"].Exec(data.UserID, data.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("DeleteUser: failed to delete user:", err)
		return
	}

	ra, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("DeleteUser: failed to convert affect rows:", err)
		return
	}

	switch ra {
	case 0:
		w.WriteHeader(http.StatusConflict)
		log.Println("DeleteUser: user does not exist")
	case 1:
		w.WriteHeader(http.StatusOK)
		log.Println("DeleteUser: deleted user successfully")
	default:
		log.Fatalln("DeleteUser: More than 1 row has been affected by DeleteUser, stopping service. USERID: ", data.UserID, "PASSWORD: ", data.Password)
	}
}

func (s *Service) CreateNewList(w http.ResponseWriter, r *http.Request) {
	// the "new" list id is coming from a json
	// need userID to make it work
}

func (s *Service) UpdateList(w http.ResponseWriter, r *http.Request) {
	// the "new" list id is coming from a json
	// need userID to make it work
}

func (s *Service) CreateHousehold(w http.ResponseWriter, r *http.Request) {
	// the householdID is created on this end
	// need userID to make it work
}

func (s *Service) UpdateHousehold(w http.ResponseWriter, r *http.Request) {
	// need to know which household to include someone under
	// need userID to make it work
}
