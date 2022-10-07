package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const httpError string = "\nGot: %v\nWant: %v\n"

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestIndex(t *testing.T) {
	app := GetApp()

	t.Run("GET on index", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response, _ := app.Test(request, 1)

		got := response.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf(httpError, got, want)
		}
	})
	t.Run("GET on notfound", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/notfound", nil)
		response, _ := app.Test(request, 1)

		got := response.StatusCode
		want := http.StatusNotFound

		if got != want {
			t.Errorf(httpError, got, want)
		}
	})
}
