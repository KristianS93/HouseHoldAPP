package service

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var ServerIns = Server{}

func TestMain(m *testing.M) {

	//Initialize server settings¢
	ServerIns.Init()
	//Run server/service instance.

	//addint test data
	// AddTestData()

	// ServerIns.Run()

	os.Exit(m.Run())
}

func TestGetList(t *testing.T) {
	// req, _ := http.NewRequest("GET", "/product?ListId=62fa8c527abec12155c907c3", nil)
	req, _ := http.NewRequest("GET", "/product?ListId=62fa8c527abec12155c907c3", nil)

	res := executeRequest(req)

	checkResCode(t, http.StatusBadRequest, res.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ServerIns.Router.ServeHTTP(rr, req)
	return rr
}

func checkResCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %dn", expected, actual)
	}
}
