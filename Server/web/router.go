package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const (
	AddItemURL    string = "http://localhost:5003/AddItem"
	ChangeItemURL string = "http://localhost:5003/ChangeItem"
)

// Routes is the Router of the server, spreading traffic to relevant handlerFuncs.
// The input taken is the given request, which is also used to call a handleFunc on.
func (s *Server) Routes(r *mux.Router) {
	// file server for css and js
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/static/css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/static/js"))))

	r.HandleFunc("/favicon.ico", s.favIcon).Methods("GET")

	// all route endpoints
	r.HandleFunc("/", s.index).Methods("GET")
	r.HandleFunc("/logout", s.logout)
	r.HandleFunc("/mealplanner", s.mealplanner)
	r.HandleFunc("/grocerylist", s.groceryList).Methods("GET")

	// form handlers
	r.HandleFunc("/changeitem", s.ChangeItem).Methods("PATCH")
	r.HandleFunc("/additem", s.additem).Methods("POST")
}

// favIcon serves the favourite icon.
func (s *Server) favIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/favicon.ico")
}

// index handles the frontpage.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s.serveSite(w, r, "index", nil)
}

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Logged out")
}

func (s *Server) mealplanner(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Meal Planner")
}

func (s *Server) groceryList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	xi := []Item{
		{
			ItemID:   "xd",
			ItemName: "pølse",
			Quantity: "3",
			Unit:     "stk",
		},
		{
			ItemID:   "notthisone",
			ItemName: "fladskærm",
			Quantity: "42",
			Unit:     "paller",
		},
	}

	tpl := TmplData{
		Data:   xi,
		Errors: GetAlert(w, r),
		User: UserData{
			Name:     "Krath",
			LoggedIn: true,
		},
	}

	s.serveSite(w, r, "grocerylist", tpl)
}

func (s *Server) additem(w http.ResponseWriter, r *http.Request) {
	item := struct {
		ListID   string
		ItemName string
		Quantity string
		Unit     string
	}{
		ListID:   "",
		ItemName: strings.ToLower(r.FormValue("name")),
		Quantity: strings.ToLower(r.FormValue("quantity")),
		Unit:     strings.ToLower(r.FormValue("unit")),
	}

	if item.ItemName == "" || item.Quantity == "" || item.Unit == "" {
		addAlert(w, r, Warning, "One field was empty, please fill all fields appropriately.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	mi, err := json.Marshal(item)
	if err != nil {
		fmt.Println("marshal went wrong")
	}

	// crashes program when external microservice is offline
	// maybe fix with http.Client
	res, err := http.Post(AddItemURL, "application/json", bytes.NewBuffer(mi))
	if err != nil {
		log.Println("request to additem failed", err)
		addAlert(w, r, Danger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	if res.StatusCode != http.StatusOK {
		addAlert(w, r, Danger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
}

func (s *Server) ChangeItem(w http.ResponseWriter, r *http.Request) {
	type changeItem struct {
		Id       string `json:"Id"`
		Name     string `json:"Name"`
		Quantity string `json:"Quantity"`
		Unit     string `json:"Unit"`
	}
	var uItem changeItem

	err := json.NewDecoder(r.Body).Decode(&uItem)
	if err != nil {
		log.Println(err)
	}

	if uItem.Id == "" || uItem.Name == "" || uItem.Quantity == "" || uItem.Unit == "" {
		addAlert(w, r, Warning, "One field was empty, please fill all fields appropriately.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(uItem.Quantity)
	if err != nil {
		addAlert(w, r, Danger, "Internal error.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if n < 1 {
		addAlert(w, r, Danger, "Quantity must be at least 1.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mi, err := json.Marshal(uItem)
	if err != nil {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("changeItem failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPatch, ChangeItemURL, bytes.NewBuffer(mi))
	if err != nil {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("changeItem newRequest: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("changeItem DO: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.StatusCode != http.StatusOK {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("changeItem unexpected status code")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
