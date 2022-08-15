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

	testcasesStatusCode := []GetListTests{
		{"Correct request should output 200", ExtraData{"GET", "/GetList?ListId=62fa8c527abec12155c907c3"}, http.StatusOK},
		{"Wrong listId should output 400", ExtraData{"GET", "/GetList?ListId=62fa8c527abec12155c90c3"}, http.StatusBadRequest},
		{"Wrong request method should output 405", ExtraData{"POST", "/GetList?ListId=62fa8c527abec12155c907c3"}, http.StatusMethodNotAllowed},
		{"Wrong url param should output 400", ExtraData{"GET", "/GetList?Id=62fa8c527abec12155c907c3"}, http.StatusBadRequest},
		{"No list id provided should output 400", ExtraData{"GET", "/GetList?ListId="}, http.StatusBadRequest},
		{"No url params should output 405", ExtraData{"GET", "/GetList"}, http.StatusMethodNotAllowed},
	}

	for _, v := range testcasesStatusCode {
		req, _ := http.NewRequest(v.Input.method, v.Input.url, nil)
		res := executeRequest(req)
		checkResCode(t, v.Output, res.Code, v.Name)
	}

	//Test specific json:

	//DENNE TEST ER EGENTLIG RIGTIG NOK MEN SOMEHOW DØR DEN I COMPARE, STRENGE TJEKKET I ANDET PROGRAM.
	req, _ := http.NewRequest("GET", "/GetList?ListId=62fa8c527abec12155c907c3", nil)
	res := executeRequest(req)

	expectedStr := `{"Succes":"List Retrieved","Items":[{"ID":"62faac9cad635f2233f86be3","ItemName":"Test item1","Quantity":"4","Unit":"pakker"},{"ID":"62faac9cad635f2233f86be4","ItemName":"Test item2","Quantity":"5","Unit":"stk"}]}`
	if body := res.Body.String(); body != expectedStr {
		t.Errorf("Testing succesfull test json data. Got %s", body)
	}

	//Delete data from test

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
