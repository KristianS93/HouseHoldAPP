package web

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"
	"time"
	"unicode"

	"github.com/google/uuid"
)

func (s *Server) serveSite(w http.ResponseWriter, r *http.Request, tplName string, data interface{}) {
	err := s.templateGet(tplName).ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Println("errServe: ", err)
	}
}

const (
	templatesBasePath = "templates/"
	templatesExt      = ".gohtml"
)

func (s *Server) parseTemplate(name, path string) {
	if path == "" {
		path = name
	}

	if _, ok := s.Templates[name]; ok {
		log.Panic("template", name)
		return
	}

	s.Templates[name] = template.Must(template.ParseFiles(templatesBasePath+"base"+templatesExt, templatesBasePath+name+templatesExt))
}

func (s *Server) RenewSession(w http.ResponseWriter, c *http.Cookie) {
	v := c.Value
	c.MaxAge = -1

	if ses, ok := s.Sessions[v]; ok {
		ses.LastActivity = time.Now()
		s.Sessions[v] = ses
		c.MaxAge = SessionTimeOut
	}
	http.SetCookie(w, c)
}

// StartSession begins a session on the server, generating a UUID cookie value,
// should be called after a login has been verified.
func (s *Server) StartSession(w http.ResponseWriter, r *http.Request, listid, householdid, name string) {
	id := uuid.New()
	ses := Session{
		LastActivity: time.Now(),
		ListID:       listid,
		HouseHoldID:  householdid,
		Name:         name,
	}
	s.Sessions[id.String()] = ses

	c := http.Cookie{
		Name:   "session",
		Value:  id.String(),
		MaxAge: SessionTimeOut,
	}
	http.SetCookie(w, &c)
}

func (s *Server) templateGet(name string) *template.Template {
	if _, ok := s.Templates[name]; !ok {
		return s.Templates["404.tmpl"]
	}

	return s.Templates[name]
}

func validLogin(email, password string) bool {
	if validEmail(email) && validPassword(password) {
		return true
	}
	return false
}

// validPassword validates if the input password
// upholds the criteria specified in the function.
// By default the password must contain:
//
// - 1 uppercase letter
//
// - 1 lowercase letter
//
// - 1 number
//
// - 1 special character
//
// And have a length between 8 and 32, inclusive.
func validPassword(p string) bool {
	var (
		hasLength  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(p) >= 8 && len(p) <= 32 {
		hasLength = true
	} else {
		return false
	}

	for _, char := range p {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		default:
			return false
		}
	}

	return hasLength && hasUpper && hasLower && hasNumber && hasSpecial
}

// validEmail validates whether the input email
// is a valid email or not, specified by a Regex string.
func validEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
