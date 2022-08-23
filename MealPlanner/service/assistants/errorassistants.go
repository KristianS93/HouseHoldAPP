package assistants

import (
	"io"
	"net/http"
)

// WrongMethod Check the method from the r pointer and the submitted method, returns a booleans, if true the method is WRONG
func WrongMethod(r *http.Request, w *http.ResponseWriter, method string) bool {
	returnVal := false
	if r.Method != method {
		(*w).WriteHeader(405)
		io.WriteString((*w), `{"Error": "Wrong method"}`)
		returnVal = true
	}
	return returnVal
}
