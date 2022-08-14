package service

import (
	"net/http"
	"os"
	"testing"
)

var TestServer = Server{}

func TestMain(m *testing.M) {
	TestServer.Init()
	TestServer.Run()

	os.Exit(m.Run())
}

func TestGetlist(t *testing.T) {

	type CorrectInput struct {
		Name   string
		Input  string
		Output string
	}

	correct := []CorrectInput{{
		"Correct input", "list id = 234342", "Succes",
	},
	}

	var w http.ResponseWriter
	var r *http.Request

	for _, v := range correct {
		TestServer.GetList(w, r)
		w.

	}

}
