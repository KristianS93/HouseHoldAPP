package web

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

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
