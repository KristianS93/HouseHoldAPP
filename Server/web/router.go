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
	// maybe new name
	UpdateListURL string = "http://localhost:5003/UpdateList"
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
	r.HandleFunc("/updatelist", s.updatelist).Methods("POST")
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
		Errors: getAlert(w, r),
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

	// fix, bad approach
	type PayLoad struct {
		Success string `json:"Succes"`
		Error   string `json:"Error"`
	}
	var p PayLoad
	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		log.Println("failed to decode")
		addAlert(w, r, Danger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	if p.Success != "" {
		addAlert(w, r, Success, "Item was successfully added to Grocery List.")
	} else {
		addAlert(w, r, Danger, "Internal error.")
	}

	http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
}

func (s *Server) updatelist(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// refactor to make a handler for this, way too long and repetitive
		log.Println("updatelist parseform: ", err)
		addAlert(w, r, Danger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}
	n := r.PostForm["name"]
	q := r.PostForm["quantity"]
	u := r.PostForm["unit"]
	id := r.PostForm["itemid"]

	updatedItems := make([]Item, len(id))
	for i := 0; i < len(id); i++ {
		updatedItems[i] = Item{
			ItemName: n[i],
			Quantity: q[i],
			Unit:     u[i],
			ItemID:   id[i],
		}
	}

	// get list id from cookie and session
	list := struct {
		ListID string
		Items  []Item
	}{
		ListID: "hulahoop",
		Items:  updatedItems,
	}

	ml, err := json.Marshal(list)
	if err != nil {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("updatelist failed to marshal: ", err)
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPatch, UpdateListURL, bytes.NewBuffer(ml))
	log.Println("errRequest", err)
	if err != nil {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("updatelist newRequest: ", err)
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("updatelist DO: ", err)
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	if res.StatusCode != http.StatusOK {
		addAlert(w, r, Danger, "Internal error.")
		log.Println("updatelist unexpected status code")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	// type PayLoad struct{}
	// err = json.NewDecoder(res.Body).Decode(&p)
	// if err != nil {
	// 	addAlert(w, r, Danger, "Internal error.")
	// 	log.Println("updatelist DO: ", err)
	// 	http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
	// 	return
	// }
}
