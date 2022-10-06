package web

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"
// )

// const stdError string = "\nGot: %v\nWant: %v\nGiven: %v\n"
// const httpError string = "\nGot: %v\nWant: %v\n"

// var server Server

// func TestMain(m *testing.M) {
// 	server = Server{}
// 	// server.Router = mux.NewRouter()
// 	// server.Routes(server.Router)
// 	server.Init()

// 	code := m.Run()
// 	os.Exit(code)
// }

// func TestIndex(t *testing.T) {
// 	t.Run("get request", func(t *testing.T) {
// 		request, err := http.NewRequest(http.MethodGet, "/", nil)
// 		if err != nil {
// 			t.Error("failed to build NewRequest:", err)
// 		}
// 		response := httptest.NewRecorder()
// 		// log.Println(request)
// 		// log.Println(response)
// 		server.Router.ServeHTTP(response, request)
// 		got := response.Result().StatusCode
// 		want := http.StatusOK

// 		if got != want {
// 			t.Errorf(httpError, got, want)
// 		}
// 	})
// }
