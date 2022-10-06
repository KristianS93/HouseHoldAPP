package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"server/web/validation"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"server/web/validation"
// 	"strconv"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// 	"golang.org/x/crypto/bcrypt"
// )

const (
	GroceryListAddItem    string = "http://localhost:5003/AddItem"
	GroceryListChangeItem string = "http://localhost:5003/ChangeItem"
	GroceryListGetList    string = "http://localhost:5003/GetList"
	GroceryListClearList  string = "http://localhost:5003/ClearList"
	UserSystemLogin       string = "http://localhost:5001/Login"
	UserSystemCreateUser  string = "http://localhost:5001/CreateUser"
)

type NavbarData struct {
	LoggedIn bool
	Name     string
}

// // Routes is the Router of the server, spreading traffic to relevant handlerFuncs.
// // The input taken is the given request, which is also used to call a handleFunc on.
// // func (s *Server) Routes(r *mux.Router) {
// // 	// file server for css and js
// // 	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/static/css"))))
// // 	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/static/js"))))

// // 	r.HandleFunc("/favicon.ico", s.favIcon).Methods(http.MethodGet)

// // 	// site endpoints
// // 	r.HandleFunc("/", s.index).Methods(http.MethodGet)
// // 	r.HandleFunc("/logout", s.logout)
// // 	// should be protected by middleware
// // 	r.HandleFunc("/mealplanner", s.mealplanner)
// // 	r.HandleFunc("/grocerylist", s.groceryList).Methods(http.MethodGet)

// // 	// form handlers or similar non-site endpoints
// // 	r.HandleFunc("/changeitem", s.ChangeItem).Methods(http.MethodPatch)
// // 	r.HandleFunc("/additem", s.additem).Methods(http.MethodPost)
// // 	r.HandleFunc("/clearlist", s.ClearList).Methods(http.MethodGet)
// // 	r.HandleFunc("/login", s.Login).Methods(http.MethodPost)
// // 	r.HandleFunc("/register", s.Register).Methods(http.MethodPost)

// // 	// any default traffic
// // 	r.NotFoundHandler = http.HandlerFunc(handle404)
// // }

// func handle404(w http.ResponseWriter, r *http.Request) {
// 	log.Println("user found a non existent endpoint")
// 	w.Write([]byte("Hello, site not found."))
// }

// index handles the frontpage.
func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		// "NavBar": NavbarData{
		// 	LoggedIn: true,
		// 	Name:     "Krath",
		// },
	})
}

// Remember to add alerts?
func GroceryList(c *fiber.Ctx) error {
	listID := "62fd4bc950c4443769551c49"
	res, err := http.Get(fmt.Sprintf("%s?ListId=%s", GroceryListGetList, listID))
	if err != nil {
		log.Println("bad getlist: ", err)
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	type GroceryListItem struct {
		Name     string `json:"ItemName"`
		Quantity string `json:"Quantity"`
		Unit     string `json:"Unit"`
		ID       string `json:"ID"`
	}
	var Items struct {
		Items []GroceryListItem `json:"Items"`
	}

	err = json.NewDecoder(res.Body).Decode(&Items)
	if err != nil {
		log.Println("bad decode: ", err)
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	return c.Render("grocerylist", fiber.Map{
		"NavBar": NavbarData{
			LoggedIn: true,
			Name:     "Krath",
		},
		"Items": Items.Items,
	})
}

// func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Logged out")
// }

// func (s *Server) mealplanner(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Meal Planner")
// }

// Remember to add alerts
func Additem(c *fiber.Ctx) error {
	item := Item{
		ListId:   "62fd4bc950c4443769551c49",
		ItemName: strings.ToLower(c.FormValue("name")),
		Quantity: strings.ToLower(c.FormValue("quantity")),
		Unit:     strings.ToLower(c.FormValue("unit")),
	}

	if item.ItemName == "" || item.Quantity == "" || item.Unit == "" {
		return c.Redirect("/grocerylist", fiber.StatusSeeOther)
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
		return c.Redirect("/grocerylist", fiber.StatusSeeOther)
	}

	if res.StatusCode != http.StatusOK {
		log.Println("additem wrong statuscode")
		return c.Redirect("grocerylist", fiber.StatusSeeOther)
	}

	return c.Redirect("grocerylist", fiber.StatusSeeOther)
}

func ChangeItem(c *fiber.Ctx) error {
	type changeItem struct {
		Id       string `json:"Id"`
		ItemName string `json:"ItemName"`
		Quantity string `json:"Quantity"`
		Unit     string `json:"Unit"`
	}
	var uItem changeItem

	err := json.Unmarshal(c.Body(), &uItem)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// stripping leading a from html element
	// leading a is important due to document.querySelectorAll
	// not accepting id's with leading digits
	uItem.Id = uItem.Id[1:]

	if uItem.Id == "" || uItem.ItemName == "" || uItem.Quantity == "" || uItem.Unit == "" {
		//addAlert(w, Warning, "One field was empty, please fill all fields appropriately.")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	n, err := strconv.Atoi(uItem.Quantity)
	if err != nil {
		// addAlert(w, Danger, "Internal error.")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if n < 1 {
		// addAlert(w, Danger, "Quantity must be at least 1.")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	mi, err := json.Marshal(uItem)
	if err != nil {
		// addAlert(w, Danger, "Internal error.")
		log.Println("changeItem failed to marshal: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	req, err := http.NewRequest(http.MethodPatch, GroceryListChangeItem, bytes.NewBuffer(mi))
	if err != nil {
		// addAlert(w, Danger, "Internal error.")
		log.Println("changeItem newRequest: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		// addAlert(w, Danger, "Internal error.")
		log.Println("changeItem DO: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.StatusCode != http.StatusOK {
		// addAlert(w, Danger, "Internal error.")
		log.Println("changeItem unexpected status code")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func ClearList(c *fiber.Ctx) error {
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
		// AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Marshal gone bad.", err))
		return c.Redirect("/grocerylist", fiber.StatusSeeOther)
	}

	nr, err := http.NewRequest(http.MethodDelete, GroceryListClearList, bytes.NewBuffer(ml))
	if err != nil {
		// AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Issue creating new request.", err))
		return c.Redirect("grocerylist", fiber.StatusSeeOther)
	}

	client := http.Client{}
	res, err := client.Do(nr)
	if err != nil {
		// AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Issue performing new request.", err))
		return err
	}

	if res.StatusCode != http.StatusOK {
		// AlertLog(w, Danger, InternalError, fmt.Sprint("ClearList: Wrong StatusCode in response.", err))
		return errors.New("clearlist: unexpected status code")
	}

	// addAlert(w, Success, ListCleared)
	return c.Redirect("grocerylist", fiber.StatusSeeOther)
}

func Login(c *fiber.Ctx) error {
	type login struct {
		UserID   string `json:"UserID"`
		Password string `json:"Password"`
	}
	var user login

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if errs := validation.CheckEmail(user.UserID); len(errs) != 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errs := validation.CheckPassword(user.Password); len(errs) != 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	tempEmail, err := bcrypt.GenerateFromPassword([]byte(user.UserID), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tempPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	user.UserID = string(tempEmail)
	user.Password = string(tempPassword)

	mu, err := json.Marshal(user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	res, err := http.Post(UserSystemLogin, "application/json", bytes.NewBuffer(mu))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusInternalServerError:
			return c.SendStatus(fiber.StatusInternalServerError)
		case http.StatusNotFound:
			return c.SendStatus(fiber.StatusNotFound)
		default:
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	// type sesInfo struct {
	// 	FirstName     string `json:"FirstName"`
	// 	GroceryListID string `json:"ListID"`
	// 	HouseholdID   string `json:"HouseholdID"`
	// }
	// var ns sesInfo

	// json.NewDecoder(r.Body).Decode(&ns)

	// s.StartSession(w, r, ns.GroceryListID, ns.HouseholdID, ns.FirstName)
	// addAlert(w, Success, LoginSuccess)
	return c.SendStatus(fiber.StatusOK)
}

// func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
// 	// was just testing, but this works with the front end
// 	type errors struct {
// 		Errors []string `json:"Errors"`
// 	}
// 	e := errors{
// 		Errors: []string{"Testing", "Testing123", "still testing"},
// 	}
// 	w.WriteHeader(http.StatusBadRequest)
// 	json.NewEncoder(w).Encode(e)
// }
