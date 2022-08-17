package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grocerylist/database"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var ServerIns = Server{}
var ListClient = database.MongClient{}

func TestMain(m *testing.M) {

	//Initialize server settings¢
	ServerIns.Init()

	//Db connection
	//Create test DB items collection.
	//addint test data
	AddTestData()

	code := m.Run()
	DeleteTestData()
	//clearData()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ServerIns.Router.ServeHTTP(rr, req)
	return rr
}

func checkResCode(t *testing.T, expected, actual int, testname string) {
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

func TestCreateList(t *testing.T) {

	type CreateListTests teststruct

	testcasesStatusCode := []CreateListTests{
		{"Correct request should output 200", ExtraData{"POST", "/CreateList", `{"HouseholdId": "TestId"}`}, http.StatusOK},
		{"Wrong method should return 405", ExtraData{"GET", "/CreateList", `{"HouseholdId": "TestId"}`}, http.StatusMethodNotAllowed},
		{"Wrong json should return 400", ExtraData{"POST", "/CreateList", `"HouseholdId": "TestId"}`}, http.StatusBadRequest},
		{"Correct request no household id should return 400", ExtraData{"POST", "/CreateList", `{"HouseholdId": ""}`}, http.StatusBadRequest},
		{"Correct request with already existing household id should return 400", ExtraData{"POST", "/CreateList", `{"HouseholdId": "TestId"}`}, http.StatusBadRequest},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, CreateReader(v.Input.body))
		res := executeRequest(req)
		checkResCode(t, v.Output, res.Code, v.Name)
	}

	fmt.Println("Passed CreateList Tests")
}

func TestGetList(t *testing.T) {

	type GetListTests teststruct

	testcasesStatusCode := []GetListTests{
		{"Correct request should output 200", ExtraData{"GET", "/GetList?ListId=62fa8c527abec12155c907c3", ""}, http.StatusOK},
		{"Wrong listId should output 400", ExtraData{"GET", "/GetList?ListId=62fa8c527abec12155c90c3", ""}, http.StatusBadRequest},
		{"Wrong request method should output 405", ExtraData{"POST", "/GetList?ListId=62fa8c527abec12155c907c3", ""}, http.StatusMethodNotAllowed},
		{"Wrong url param should output 400", ExtraData{"GET", "/GetList?Id=62fa8c527abec12155c907c3", ""}, http.StatusBadRequest},
		{"No list id provided should output 400", ExtraData{"GET", "/GetList?ListId=", ""}, http.StatusBadRequest},
		{"No url params should output 405", ExtraData{"GET", "/GetList", ""}, http.StatusMethodNotAllowed},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, nil)
		res := executeRequest(req)
		checkResCode(t, v.Output, res.Code, v.Name)
	}

	//Test specific json:
	req, _ := http.NewRequest("GET", "/GetList?ListId=62fa8c527abec12155c907c3", nil)
	res := executeRequest(req)

	var di ItemHolder
	err := json.NewDecoder(res.Body).Decode(&di)
	if err != nil {
		t.Errorf("Test: %s - ", "Getting specific json")
	}

	if di.Succes != "List Retrieved" {
		t.Errorf("Test: %s - ", "Getting specific json")
	}
	fmt.Println("Passed GetList Tests")
}

func TestClearList(t *testing.T) {
	type ClearListTest teststruct

	testcasesStatusCode := []ClearListTest{
		{"Correct ClearList request should return 200", ExtraData{"DELETE", "/ClearList", `{"ListId": "62fa8c527abec12155c907c3"}`}, http.StatusOK},
		{"Wrong request method should return 405", ExtraData{"GET", "/ClearList", `{"ListId": "62fa8c527abec12155c907c3"}`}, http.StatusMethodNotAllowed},
		{"Wrong json in request should return 400", ExtraData{"DELETE", "/ClearList", `{}`}, http.StatusBadRequest},
		{"No list id in request should return 400", ExtraData{"DELETE", "/ClearList", `{"ListId": ""}`}, http.StatusBadRequest},
		{"List id doesnt exist should return 400", ExtraData{"DELETE", "/ClearList", `{"ListId": "XXXXXXXXX"}`}, http.StatusBadRequest},
		{"List is empty should return 400", ExtraData{"DELETE", "/ClearList", `{"ListId": "62fa8c527abec12155c907c3"}`}, http.StatusBadRequest},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, CreateReader(v.Input.body))
		res := executeRequest(req)
		checkResCode(t, v.Output, res.Code, v.Name)
	}
	fmt.Println("Passed ResetList Tests")
}

func TestDeleteList(t *testing.T) {
	type DeleteListTest teststruct

	testcasesStatusCode := []DeleteListTest{
		{"Correct DeleteList request should return 200", ExtraData{"DELETE", "/DeleteList", `{"ListId": "62fa8c527abec12155c907c3"}`}, http.StatusOK},
		{"Wrong method for DeleteList request should return 405", ExtraData{"GET", "/DeleteList", `{"ListId": "62fa8c527abec12155c907c3"}`}, http.StatusMethodNotAllowed},
		{"No json for DeleteList request should return 400", ExtraData{"DELETE", "/DeleteList", `{}`}, http.StatusBadRequest},
		{"List id empty for DeleteList request should return 400", ExtraData{"DELETE", "/DeleteList", `{"ListId": ""}`}, http.StatusBadRequest},
		{"Non existing List id for DeleteList request should return 400", ExtraData{"DELETE", "/DeleteList", `{"ListId": "xxxxyyyy"}`}, http.StatusBadRequest},
		{"Deleting already deleted list id request should return 400", ExtraData{"DELETE", "/DeleteList", `{"ListId": "62fa8c527abec12155c907c3"}`}, http.StatusBadRequest},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, CreateReader(v.Input.body))
		res := executeRequest(req)
		checkResCode(t, v.Output, res.Code, v.Name)
	}
	fmt.Println("Passed DeleteList Tests")
}
