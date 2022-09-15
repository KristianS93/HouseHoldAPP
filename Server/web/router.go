package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/web/validation"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

const (
	GroceryListAddItem    string = "http://localhost:5003/AddItem"
	GroceryListChangeItem string = "http://localhost:5003/ChangeItem"
	GroceryListGetList    string = "http://localhost:5003/GetList"
	GroceryListClearList  string = "http://localhost:5003/ClearList"
	UserSystemLogin       string = "http://localhost:5001/Login"
	UserSystemCreateUser  string = "http://localhost:5001/CreateUser"
)

// Routes is the Router of the server, spreading traffic to relevant handlerFuncs.
// The input taken is the given request, which is also used to call a handleFunc on.
func (s *Server) Routes(r *mux.Router) {
	// file server for css and js
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/static/css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/static/js"))))

	r.HandleFunc("/favicon.ico", s.favIcon).Methods(http.MethodGet)

	// site endpoints
	r.HandleFunc("/", s.index).Methods(http.MethodGet)
	r.HandleFunc("/logout", s.logout)
	// should be protected by middleware
	r.HandleFunc("/mealplanner", s.mealplanner)
	r.HandleFunc("/grocerylist", s.groceryList).Methods(http.MethodGet)

	// form handlers or similar non-site endpoints
	r.HandleFunc("/changeitem", s.ChangeItem).Methods(http.MethodPatch)
	r.HandleFunc("/additem", s.additem).Methods(http.MethodPost)
	r.HandleFunc("/clearlist", s.ClearList).Methods(http.MethodGet)
	r.HandleFunc("/login", s.Login).Methods(http.MethodPost)
	r.HandleFunc("/register", s.Register).Methods(http.MethodPost)
}

// favIcon serves the favourite icon.
func (s *Server) favIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/favicon.ico")
}

// index handles the frontpage.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	// remember to get and add alerts

	// these are dummy values, should possibly be middleware
	// tpl := TmplData{
	// 	Errors: GetAlert(w, r),
	// }

	s.serveSite(w, r, "index", nil)
}

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Logged out")
}

func (s *Server) mealplanner(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Meal Planner")
}

func (s *Server) groceryList(w http.ResponseWriter, r *http.Request) {
	listID := "62fd4bc950c4443769551c49"
	res, err := http.Get(fmt.Sprintf("%s?ListId=%s", GroceryListGetList, listID))
	if err != nil {
		addAlert(w, Danger, "Internal error.")
		log.Println("bad getlist: ", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	type GetList struct {
		Success string `json:"Success"`
		XI      []Item `json:"Items"`
	}

	var ItemHolder GetList

	err = json.NewDecoder(res.Body).Decode(&ItemHolder)
	if err != nil {
		addAlert(w, Danger, "Internal error.")
		log.Println("bad decode: ", err)
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	tpl := TmplData{
		Data:   ItemHolder.XI,
		Errors: GetAlert(w, r),
		User: UserData{
			Name:     "Krath",
			LoggedIn: true,
		},
	}

	s.serveSite(w, r, "grocerylist", tpl)
}

func (s *Server) additem(w http.ResponseWriter, r *http.Request) {
	item := Item{
		ListId:   "62fd4bc950c4443769551c49",
		ItemName: strings.ToLower(r.FormValue("name")),
		Quantity: strings.ToLower(r.FormValue("quantity")),
		Unit:     strings.ToLower(r.FormValue("unit")),
	}

	if item.ItemName == "" || item.Quantity == "" || item.Unit == "" {
		addAlert(w, Warning, "One field was empty, please fill all fields appropriately.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	xi := []Item{item}

	mi, err := json.Marshal(xi)
	if err != nil {
		fmt.Println("marshal went wrong")
	}

	// crashes program when external microservice is offline
	// maybe fix with http.Client
	res, err := http.Post(GroceryListAddItem, "application/json", bytes.NewBuffer(mi))
	if err != nil {
		log.Println("request to additem failed", err)
		addAlert(w, Danger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}
	log.Println("statuscode:", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		log.Println("additem wrong statuscode")
		addAlert(w, Danger, "Internal error.")
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	addAlert(w, Success, "Item added successfully.")
	http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
}

func (s *Server) ChangeItem(w http.ResponseWriter, r *http.Request) {
	type changeItem struct {
		Id       string `json:"Id"`
		ItemName string `json:"ItemName"`
		Quantity string `json:"Quantity"`
		Unit     string `json:"Unit"`
	}
	var uItem changeItem

	err := json.NewDecoder(r.Body).Decode(&uItem)
	if err != nil {
		log.Println(err)
	}

	// stripping leading a from html element
	// leading a is important due to document.querySelectorAll
	// not accepting id's with leading digits
	uItem.Id = uItem.Id[1:]

	if uItem.Id == "" || uItem.ItemName == "" || uItem.Quantity == "" || uItem.Unit == "" {
		addAlert(w, Warning, "One field was empty, please fill all fields appropriately.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(uItem.Quantity)
	if err != nil {
		addAlert(w, Danger, "Internal error.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if n < 1 {
		addAlert(w, Danger, "Quantity must be at least 1.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mi, err := json.Marshal(uItem)
	if err != nil {
		addAlert(w, Danger, "Internal error.")
		log.Println("changeItem failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPatch, GroceryListChangeItem, bytes.NewBuffer(mi))
	if err != nil {
		addAlert(w, Danger, "Internal error.")
		log.Println("changeItem newRequest: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		addAlert(w, Danger, "Internal error.")
		log.Println("changeItem DO: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.StatusCode != http.StatusOK {
		addAlert(w, Danger, "Internal error.")
		log.Println("changeItem unexpected status code")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) ClearList(w http.ResponseWriter, r *http.Request) {
	// the id should be retrieved from the session cookie
	// and session runtime database
	listID := "62fd4bc950c4443769551c49"
	listObj := struct {
		ListId string
	}{
		ListId: listID,
	}

	ml, err := json.Marshal(listObj)
	if err != nil {
		AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Marshal gone bad.", err))
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	nr, err := http.NewRequest(http.MethodDelete, GroceryListClearList, bytes.NewBuffer(ml))
	if err != nil {
		AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Issue creating new request.", err))
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	client := http.Client{}
	res, err := client.Do(nr)
	if err != nil {
		AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Issue performing new request.", err))
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	if res.StatusCode != http.StatusOK {
		AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Wrong StatusCode in response.", err))
		http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
		return
	}

	addAlert(w, Success, ListCleared)
	http.Redirect(w, r, "/grocerylist", http.StatusSeeOther)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	type login struct {
		UserID   string `json:"UserID"`
		Password string `json:"Password"`
	}
	var user login

	json.NewDecoder(r.Body).Decode(&user)

	if errs := validation.CheckEmail(user.UserID); !errs {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("login: bad email")
		return
	}

	if errs := validation.CheckPassword(user.Password); len(errs) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("login: bad password")
		return
	}

	tempEmail, err := bcrypt.GenerateFromPassword([]byte(user.UserID), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("bcrypt userid")
		return
	}
	user.UserID = string(tempEmail)

	tempPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("bcrypt password")
		return
	}
	user.Password = string(tempPassword)

	mu, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("json marshal")
		return
	}

	res, err := http.Post(UserSystemLogin, "application/json", bytes.NewBuffer(mu))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("http post")
		return
	}

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusInternalServerError:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("switch internalerror")
			return
		case http.StatusNotFound:
			w.WriteHeader(http.StatusNotFound)
			log.Println("switch not found")
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalln("unexpected status code returned from login attempt:", res.StatusCode)
			return
		}
	}

	type sesInfo struct {
		FirstName     string `json:"FirstName"`
		GroceryListID string `json:"ListID"`
		HouseholdID   string `json:"HouseholdID"`
	}
	var ns sesInfo

	json.NewDecoder(r.Body).Decode(&ns)

	s.StartSession(w, r, ns.GroceryListID, ns.HouseholdID, ns.FirstName)
	addAlert(w, Success, LoginSuccess)
	w.WriteHeader(http.StatusOK)
	log.Println("Session started")
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	// was just testing, but this works with the front end
	type errors struct {
		Errors []string `json:"Errors"`
	}
	e := errors{
		Errors: []string{"Testing", "Testing123", "still testing"},
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(e)
}
