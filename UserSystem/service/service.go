package service

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	Router   *mux.Router
	HostName string
	HostPort string
	DB       *sql.DB
	// Statements map[string]*sql.Stmt
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
}

func (s *Service) Run() {
	log.Println("Starting UserSystem on " + s.HostName + s.HostPort)
	err := http.ListenAndServe((s.HostName + s.HostPort), s.Router)
	if err != nil {
		log.Fatalln("failed to listen and serve")
	}
}
