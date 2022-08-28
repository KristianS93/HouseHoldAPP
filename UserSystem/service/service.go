package service

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	Router     *mux.Router
	HostName   string
	HostPort   string
	DB         *sql.DB
	Statements map[string]*sql.Stmt
}

func (s *Service) Init() {
	if s.Router == nil {
		s.Router = mux.NewRouter()
		s.Routes(s.Router)
	}

	if s.HostName == "" {
		s.HostName = "localhost"
	}

	if s.HostPort == "" {
		s.HostPort = ":5001"
	}

	if s.DB == nil {
		var err error
		s.DB, err = sql.Open("sqlite3", "database/storage/usersystem.sqlite")
		if err != nil {
			log.Fatalln("failed to open database")
		}
	}

	if s.Statements == nil {
		s.Statements = make(map[string]*sql.Stmt)

		// Maybe refactor into constants and prepare, close etc. in functions
		ns := []NewStatements{
			{
				"SELECT firstName, listID, householdID FROM USERS WHERE userID = $1 AND password = $2",
				"Login",
			},
			{
				"INSERT INTO USERS (userID, password, firstName, lastName, listID, householdID) VALUES ($1, $2, $3, $4, $5, $6)",
				"CreateUser",
			},
			{
				"DELETE FROM USERS WHERE userID = $1 AND password = $2",
				"DeleteUser",
			},
			{
				"UPDATE USERS SET householdID = $1 WHERE userID = $2",
				"HouseHold",
			},
			{
				"SELECT householdID FROM USERS WHERE userID = $1",
				"GetHHID",
			},
			{
				"UPDATE USERS SET listID = $1 WHERE userID = $2",
				"GroceryList",
			},
		}
		// the following contains all prepared statements for later execution

		for _, v := range ns {
			stmt, err := s.DB.Prepare(v.Statement)
			if err != nil {
				log.Fatalf("Failed to prepare %s query statement, err: %s", v.Identifier, err)
			}
			s.Statements[v.Identifier] = stmt
		}
	}
}

func (s *Service) Run() {
	log.Println("Starting UserSystem on " + s.HostName + s.HostPort)
	err := http.ListenAndServe((s.HostName + s.HostPort), s.Router)
	if err != nil {
		log.Fatalln("failed to listen and serve")
	}
}

type NewStatements struct {
	Statement  string
	Identifier string
}
