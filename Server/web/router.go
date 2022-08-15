package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	AddItemURL string = "http://localhost:5003/AddItem"
)

// Routes is the Router of the server, spreading traffic to relevant handlerFuncs.
// The input taken is the given request, which is also used to call a handleFunc on.
func (s *Server) Routes(r *mux.Router) {
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/static/css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/static/js"))))

	r.HandleFunc("/favicon.ico", s.favIcon)

	r.HandleFunc("/", s.index).Methods("GET")
	r.HandleFunc("/logout", s.logOut)
	r.HandleFunc("/mealplanner", s.mealPlanner)
	r.HandleFunc("/grocerylist", s.groceryList)

	r.HandleFunc("/updatelist", s.updatelist).Methods("POST")
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

	xi := []Item{
		Item{
			ItemID:   "nigger",
			ItemName: "pølse",
			Quantity: "3",
			Unit:     "stk",
		},
		Item{
			ItemID:   "black ones",
			ItemName: "fladskærm",
			Quantity: "42",
			Unit:     "paller",
		},
	}

	tpl := TmplData{
		Data:   xi,
		Errors: getAlert(w, r),
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
		ItemID:   "",
		ListID:   "62f6e4c5682d11242ec29b26",
		ItemName: strings.ToLower(r.FormValue("name")),
		Quantity: strings.ToLower(r.FormValue("quantity")),
		Unit:     strings.ToLower(r.FormValue("unit")),
	}

	if i.ItemName == "" || i.Quantity == "" || i.Unit == "" {
		addAlert(w, r, alertLevelWarning, "One field was empty, please fill all fields appropriately.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	mi, err := json.Marshal(i)
	if err != nil {
		fmt.Println("marshal went wrong")
	}

	// crashes program when external microservice is offline
	// maybe fix with http.Client
	res, err := http.Post(AddItemURL, "application/json", bytes.NewBuffer(mi))
	if err != nil {
		log.Println("request to additem failed", err)
		addAlert(w, r, alertLevelDanger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	type PayLoad struct {
		Success string `json:"Succes"`
		Error   string `json:"Error"`
	}
	var p PayLoad
	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		log.Println("failed to decode")
		addAlert(w, r, alertLevelDanger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	if p.Success != "" {
		addAlert(w, r, alertLevelSuccess, "Item was successfully added to Grocery List.")
	} else {
		addAlert(w, r, alertLevelDanger, "Internal error.")
	}

	http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
}

func (s *Server) updatelist(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	log.Println(r.PostForm["name"])
}
