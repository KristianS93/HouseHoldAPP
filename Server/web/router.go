package web

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Routes is the Router of the server, spreading traffic to relevant handlerFuncs.
// The input taken is the given request, which is also used to call a handleFunc on.
func (s *Server) Routes(r *http.ServeMux) {
	r.HandleFunc("/favicon.ico", s.favIcon)

	r.HandleFunc("/", s.index)
	r.HandleFunc("/logout", s.logOut)
	r.HandleFunc("/mealplanner", s.mealPlanner)
	r.HandleFunc("/grocerylist", s.groceryList)

	r.HandleFunc("/additem", s.addItem)
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
	s.serveSite(w, r, "index", nil)
}

func (s *Server) logOut(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Logged out")
}

func (s *Server) mealPlanner(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Meal Planner")
}

func (s *Server) groceryList(w http.ResponseWriter, r *http.Request) {
	if m := checkMethod(w, r, http.MethodGet); !m {
		return
	}
	w.WriteHeader(http.StatusOK)

	tpl := TmplData{
		Data: nil,
		Errors: []Alert{
			{
				Level:   alertLevelWarning,
				Message: "u don goofed",
			},
		},
		User: UserData{
			Name:     "Krath",
			LoggedIn: true,
		},
	}

	s.serveSite(w, r, "grocerylist", tpl)
}

func (s *Server) addItem(w http.ResponseWriter, r *http.Request) {
	if m := checkMethod(w, r, http.MethodPost); !m {
		return
	}

	i := Item{
		Name:     strings.ToLower(r.FormValue("Name")),
		Quantity: strings.ToLower(r.FormValue("Quantity")),
		Unit:     strings.ToLower(r.FormValue("Unit")),
	}
	fmt.Println(i)
	if i.Name == "" || i.Quantity == "" || i.Unit == "" {

	}
	http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)

}
