package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

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
	r.HandleFunc("/CreateHousehold", s.CreateHousehold)
}

// Login checks the LOGIN table for matching signatures
// of provided information and if appropriate login is provided,
// the user's name, listID and householdID is returned
// to the responsewriter, for the central server to handle
// and utilize in further requests and or front end display.
func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	var data login
	err := DecodeRequest(r, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
	err := DecodeRequest(r, &nu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = s.Statements["CreateUser"].Exec(nu.UserID, nu.Password, nu.FirstName, nu.LastName, nu.ListID, nu.HouseholdID)
	if err != nil {
		// following is a type assertion
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				w.WriteHeader(http.StatusConflict)
				log.Println("CreateUser: User already exists")
				return
			}
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
	err := DecodeRequest(r, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	// this is bad
	if ra != 1 {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("DeleteUser: stopping service. USERID: ", data.UserID, "PASSWORD: ", data.Password)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("DeleteUser: deleted user successfully")
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
	var id login
	err := DecodeRequest(r, &id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// The following is an example of the same functionality as above, however
	// with no side effects of the function - this would require the DecodeRequest
	// function to return the decoded json interface and the possible error, instead of only an error.
	//
	// var u login
	// temp, err := DecodeRequest(r, u)
	// if err != nil {
	// 	// check if there was an error in decoding
	// }
	// if do, ok := temp.(login); ok {
	// 	// did my type assertion work
	// }

	result, err := s.Statements["HouseHold"].Exec(uuid.New().String(), id.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateHousehold: Failed to update householdID")
		return
	}
	ra, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateHousehold: Failed to read affected rows")
		return
	}
	// this is bad
	if ra != 1 {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("CreateHousehold: Very bad, rows affected:", ra)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("CreateHousehold: HouseholdID updated succesfully")
}

// func (s *Service) UpdateHousehold(w http.ResponseWriter, r *http.Request) {
// 	type users struct {
// 		StartUser string `json:"StartUser"`
// 		DestUser  string `json:"DestUser"`
// 	}
// 	var u users
// 	err := json.NewDecoder(r.Body).Decode(&u)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Println("UpdateHousehold: Failed to decode request body: ", err)
// 		return
// 	}
// 	var hhID string
// 	err = s.Statements["GetHHID"].QueryRow(u.StartUser).Scan(&hhID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			w.WriteHeader(http.StatusNotFound)
// 			log.Println("UpdateHousehold: user not found")
// 			return
// 		}
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Println("UpdateHousehold: something wrong with DB, err:", err)
// 		return
// 	}

// 	result, err := s.Statements["HouseHold"].Exec(hhID, u.DestUser)

// 	// need to know which household to include someone under
// 	// need userID to make it work
// }
