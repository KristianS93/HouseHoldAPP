package web

import (
	"io"
	"net/http"
)

// Routes is the Router of the server, spreading traffic to relevant handlerFuncs.
// The input taken is the given request, which is also used to call a handleFunc on.
func (s *Server) Routes(r *http.ServeMux) {
	r.HandleFunc("/favicon.ico", s.favIcon)

	r.HandleFunc("/", s.index)
	r.HandleFunc("/logout", s.logOut)
	r.HandleFunc("/mealplanner", s.mealPlanner)
	r.HandleFunc("/grocerylist", s.groceryList)
}

// facIcon serves the favourite icon for the web page.
func (s *Server) favIcon(w http.ResponseWriter, r *http.Request) {
	if m := checkMethod(w, r, http.MethodGet); !m {
		return
	}
	http.ServeFile(w, r, "images/favicon.ico")
}

// index handles the frontpage of the web app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if m := checkMethod(w, r, http.MethodGet); !m {
		return
	}
	w.WriteHeader(http.StatusOK)

	// currently serves a string as the name, should serve whatever is associated with cookie value
	// or a database fetch
	s.Templates.ExecuteTemplate(w, "index.gohtml", "Andreas")
}

func (s *Server) logOut(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Logged out")
}

func (s *Server) mealPlanner(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Meal Planner")
}

func (s *Server) groceryList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Grocery List")
}
