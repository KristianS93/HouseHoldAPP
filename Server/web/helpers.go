package web

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
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

// validLogin is a wrapper for validEmail and validPassword.
//
// Returns: a bool representing the validity of the provided login;
// true if both email and password are valid, false if one of or both are invalid.
// func validLogin(email, password string) bool {
// 	return validEmail(email) && validPassword(password)
// }

// validEmail validates whether the input email
// is a valid email or not, also does an MX lookup on domain.
//
// Returns: a bool representing the validity of the provided email
// true meaning valid and false meaning invalid.
func validEmail(e string) bool {
	// splitting in prefix and domain
	xs := strings.Split(e, "@")
	if len(xs) != 2 {
		// a valid email only contains one instance of @
		return false
	}

	// checking prefix format, may contains a-z, A-Z, 0-9 and .-_ (however not consecutive)
	// anything other than this, or consecutive cases of .-_ is considered invalid
	for i, v := range xs[0] {
		switch {
		case unicode.IsUpper(v) || unicode.IsLower(v) || unicode.IsNumber(v):
			continue
		default:
			if strings.ContainsAny(string(v), ".-_") {
				// checking not out of bounds and if consecutive case of .-_
				if (i+1 <= len(xs[0])) && (strings.ContainsAny(string(xs[0][i+1]), ".-_")) {
					return false
				}
				// valid case of punctuation => skip to next rune
				continue
			}
			// only reached if invalid rune
			return false
		}
	}

	// looking up domain, returns a list of valid domains under lookup
	// err is returned when no host under specified domain exists
	// very very giga unlikely, but technically error could be
	// caused by the DNS message not being unmarshalled correctly
	_, err := net.LookupMX(xs[1])
	return err == nil

	// should technically do an email verification at this point
	// both to double check, but also to verify that the correct
	// email is associated with a given user, as they might have
	// misspelled or had a typo when entering their email
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
// - 1 special character (! @ # $ % ^ & *)
//
// And have a length between 8 and 32, inclusive.
//
// Additionally, the function validates that no invalid
// characters are present in the input password.
//
// Returns: a bool representing the validity of the provided password,
// true for a valid password and false for an invalid password.
func validPassword(p string) []string {
	var (
		hasLength    = false
		hasUpper     = false
		hasLower     = false
		hasNumber    = false
		hasSymbol    = false
		invalidChars []rune
	)

	if len(p) >= 8 && len(p) <= 32 {
		hasLength = true
	}

	for _, char := range p {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case strings.ContainsAny(string(char), "!@#$%^&*"):
			hasSymbol = true
		default:
			invalidChars = append(invalidChars, char)
		}
	}

	var warnings []string
	if !hasLength {
		warnings = append(warnings, "Password is not an appropriate length.")
	}
	if !hasUpper {
		warnings = append(warnings, "Password does not contain an upper case letter.")
	}
	if !hasLower {
		warnings = append(warnings, "Password does not contain a lower case letter.")
	}
	if !hasNumber {
		warnings = append(warnings, "Password does not contain a number.")
	}
	if !hasSymbol {
		warnings = append(warnings, "Password does not contain a symbol.")
	}
	if len(invalidChars) != 0 {
		for _, v := range invalidChars {
			warnings = append(warnings, fmt.Sprintf("Password contains an invalid character: %c", v))
		}
	}

	if len(warnings) != 0 {
		return warnings
	} else {
		return nil
	}
}
