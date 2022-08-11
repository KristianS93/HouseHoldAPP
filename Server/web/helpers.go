package web

import "net/http"

func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(405)
		w.Write([]byte("Wrong http method."))
		return true
	}
	return false
}
