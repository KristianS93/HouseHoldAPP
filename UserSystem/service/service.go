package service

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	Router   *mux.Router
	HostName string
	HostPort string
	DB       *sql.DB
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
		s.HostPort = "5001"
	}

	if s.DB == nil {
		// implement
	}

}

func (s *Service) Routes(r *mux.Router) {

}
