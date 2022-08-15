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

	//addint test data
	// AddTestData()

	code := m.Run()

	os.Exit(code)
}

func TestGetList(t *testing.T) {

	type ExtraData struct {
		method string
		url    string
	}

	type GetListTests struct {
		Name   string
		Input  ExtraData
		Output int
	}

	testcases := []GetListTests{
		{"Correct request should output 200", ExtraData{"GET", "/GetList?ListId=62fa8c527abec12155c907c3"}, http.StatusOK},
		{"Wrong listId should output 400", ExtraData{"GET", "/GetList?ListId=62fa8c527abec12155c90c3"}, http.StatusBadRequest},
		{"Wrong request method should output 405", ExtraData{"POST", "/GetList?ListId=62fa8c527abec12155c907c3"}, http.StatusMethodNotAllowed},
		{"Wrong url param should output 400", ExtraData{"GET", "/GetList?Id=62fa8c527abec12155c907c3"}, http.StatusBadRequest},
	}

	for _, v := range testcases {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, nil)
		res := executeRequest(req)
		checkResCode(t, v.Output, res.Code)
	}

	//Delete data from test

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
