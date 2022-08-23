package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var ServerIns = Server{}

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	rr.Header().Set("Access-Control-Allow-Origin", "*")
	ServerIns.Router.ServeHTTP(rr, req)
	return rr
}

func CheckResCode(t *testing.T, expected, actual int, testname string) {
	if expected != actual {
		t.Errorf("Test: %s - Expected response code %d. Got %d\n", testname, expected, actual)
	}
}

type ExtraData struct {
	method string
	url    string
	body   string
}

type teststruct struct {
	Name   string
	Input  ExtraData
	Output int
}

type reader io.Reader

func CreateReader(str string) reader {
	var jsonStr = []byte(str)
	return bytes.NewBuffer(jsonStr)
}

func TestMain(m *testing.M) {

	//Initialize server settings¢
	ServerIns.Init()

	//Db connection
	//Create test DB items collection.
	//addint test data
	// AddTestData()

	code := m.Run()
	// DeleteTestData()
	//clearData()
	os.Exit(code)
}

func TestGetMeal(t *testing.T) {

	type GetMealTest teststruct

	testcasesStatusCode := []GetMealTest{
		{"Correct GET request, output 200", ExtraData{"GET", "/GetMeal?Meal=testid", ""}, http.StatusOK},
		{"Wrong request method, output 405", ExtraData{"POST", "/GetMeal?Meal=testid", ""}, http.StatusMethodNotAllowed},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, nil)
		res := ExecuteRequest(req)
		CheckResCode(t, v.Output, res.Code, v.Name)
	}

	fmt.Println("Passed GetMeal Tests")
}
