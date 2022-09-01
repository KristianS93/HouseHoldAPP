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
	r.HandleFunc("/Login", s.Login).Methods(http.MethodPost)
	r.HandleFunc("/CreateUser", s.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/DeleteUser", s.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/CreateHousehold", s.CreateHousehold).Methods(http.MethodPost)
	r.HandleFunc("/JoinHousehold", s.JoinHousehold).Methods(http.MethodPatch)
	r.HandleFunc("/UpdateGroceryList", s.UpdateGroceryList).Methods(http.MethodPatch)
}

// Login checks the LOGIN table for matching signatures
// of provided information and if appropriate login data is provided,
// the user's name, listID and householdID is returned
// to the ResponseWriter, for the frontend to handle
// and utilize in further requests and or display purposes.
func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	var lg login
	err := DecodeRequest(r, &lg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := s.DB.Prepare("SELECT firstName, listID, householdID FROM USERS WHERE userID = $1 AND password = $2")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login: failed to prepare query to check login, err:", err)
		return
	}
	defer stmt.Close()

	var u user
	err = stmt.QueryRow(lg.UserID, lg.Password).Scan(&u.FirstName, &u.ListID, &u.HouseholdID)
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
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Login: Failed to encode response.")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	var nu NewUser
	err := DecodeRequest(r, &nu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := s.DB.Prepare("INSERT INTO USERS (userID, password, firstName, lastName, listID, householdID) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateUser: failed to prepare query to insert user, err:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(nu.UserID, nu.Password, nu.FirstName, nu.LastName, nu.ListID, nu.HouseholdID)
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

	err = CheckResult(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateUser:", err)
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

	stmt, err := s.DB.Prepare("DELETE FROM USERS WHERE userID = $1 AND password = $2")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("UpdateGroceryList: failed to prepare query to update listID, err:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.UserID, data.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("DeleteUser: failed to delete user:", err)
		return
	}

	// this should return 404 when the user is not found, ie. 0 rows were affected
	err = CheckResult(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("DeleteUser:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("DeleteUser: deleted user successfully")
}

// UpdateGroceryList updates the GroceryListID for the provided userID.
func (s *Service) UpdateGroceryList(w http.ResponseWriter, r *http.Request) {
	// incoming json contains userID and listID
	type newGroceryList struct {
		UserID string `json:"UserID"`
		ListID string `json:"ListID"`
	}
	var nl newGroceryList
	err := DecodeRequest(r, &nl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("UpdateGroceryList: error decoding json, err:", err)
		return
	}

	stmt, err := s.DB.Prepare("UPDATE USERS SET listID = $1 WHERE userID = $2")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("UpdateGroceryList: failed to prepare query to update listID, err:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(nl.ListID, nl.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("UpdateGroceryList: error executing query, err:", err)
		return
	}

	err = CheckResult(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("UpdateGroceryList:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("UpdateGroceryList: Successfully updated grocerylistID.")
}

func (s *Service) CreateHousehold(w http.ResponseWriter, r *http.Request) {
	// The following is an example of the same functionality as below, however
	// with no side effects of the function - this would require the DecodeRequest
	// function to return the decoded json interface and the possible error, instead of only an error.
	//
	// var u login
	// temp, err := DecodeRequest(r, u)
	// if err != nil {
	// 	// check if there was an error in decoding
	// }
	// if do, ok := temp.(login); ok {
	// 	// did the type assertion work
	// }

	var id login
	err := DecodeRequest(r, &id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := s.DB.Prepare("UPDATE USERS SET householdID = $1 WHERE userID = $2")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateHousehold: failed to prepare query to set household, err:", err)
		return
	}
	defer stmt.Close()

	newUUID := uuid.New()
	result, err := stmt.Exec(newUUID.String(), id.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateHousehold: Failed to update householdID")
		return
	}

	err = CheckResult(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateHousehold:", err)
		return
	}

	// automatically writes statuscode 200
	_, err = w.Write([]byte(newUUID.String()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateHousehold: failed to write newUUID to response, err:", err)
		return
	}
	log.Println("CreateHousehold: HouseholdID updated succesfully")
}

func (s *Service) JoinHousehold(w http.ResponseWriter, r *http.Request) {
	type updateHousehold struct {
		StartUser string `json:"StartUser"`
		DestUser  string `json:"DestUser"`
	}
	var u updateHousehold
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("JoinHousehold: Failed to decode request body: ", err)
		return
	}

	var hhID, glID string
	stmt, err := s.DB.Prepare("SELECT householdID, listID FROM USERS WHERE userID = $1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("JoinHousehold: failed to prepare query to get hhID and glID, err:", err)
		return
	}

	err = stmt.QueryRow(u.StartUser).Scan(&hhID, &glID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			log.Println("JoinHousehold: user not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("JoinHousehold: something wrong with DB, err:", err)
		return
	}

	stmt, err = s.DB.Prepare("UPDATE USERS SET householdID = $1, listID = $2 WHERE userID = $3")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("JoinHousehold: failed to prepare query to update hhID and glID, err:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(hhID, glID, u.DestUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("JoinHousehold: Failed to access database, err:", err)
		return
	}

	err = CheckResult(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("JoinHousehold:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("JoinHousehold: Successfully updated HouseholdID.")
}
