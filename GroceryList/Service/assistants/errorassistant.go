package assistants

import (
	"io"
	"net/http"
)

func WrongMethod(r *http.Request, w *http.ResponseWriter, method string) bool {
	returnVal := false
	if r.Method != method {
		(*w).WriteHeader(405)
		io.WriteString((*w), `{"Error": "Wrong method"}`)
		returnVal = true
	}
	return returnVal
}
