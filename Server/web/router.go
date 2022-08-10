package web

import "net/http"

// Routes is the Router of the server, spreading traffic to relevant handlerFuncs.
// The input taken is the given request, which is also used to call a handleFunc on.
func (s *Server) Routes(r *http.ServeMux) {
	r.HandleFunc("/index", s.index)
}

// index handles the frontpage of the web app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Wrong http method."))
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html")
	s.Templates.ExecuteTemplate(w, "index.gohtml", nil)
}
