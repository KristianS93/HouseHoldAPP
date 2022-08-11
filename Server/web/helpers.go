package web

import "net/http"

// checkMethod returns true when http methods are the same, and false when not.
func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong http method."))
		return false
	}
	return true
}
