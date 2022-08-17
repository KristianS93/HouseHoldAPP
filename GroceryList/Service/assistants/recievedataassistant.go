package assistants

import (
	"encoding/json"
	"io"
	"net/http"
)

func CheckData[T any](w *http.ResponseWriter, r *http.Request, items *[]T) bool {
	returnVal := false
	err := json.NewDecoder(r.Body).Decode(items)
	if err != nil {
		(*w).WriteHeader(400)
		io.WriteString((*w), `{"Error": "Bad request: Getting data"}`)
		returnVal = true
	}
	return returnVal
}
