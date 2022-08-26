package assistants

import (
	"net/http"
)

// WrongMethod Check the method from the r pointer and the submitted method, returns a booleans, if true the method is WRONG
func WrongMethod(r *http.Request, method string) bool {
	returnVal := false
	if r.Method != method {
		returnVal = true
	}
	return returnVal
}
