package web

import (
	"fmt"
	"net/http"
	"text/template"
)

type alertLevel uint8

const (
	alertLevelSuccess alertLevel = iota
	alertLevelWarning
	alertLevelDanger
)

type Alert struct {
	Level   alertLevel
	Message string
}

// checkMethod returns true when http methods are the same, and false when not.
func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong http method."))
		return false
	}
	return true
}

func (s *Server) serveSite(w http.ResponseWriter, r *http.Request, tplName string, data interface{}) {
	tpl := TmplData{
		Data: data,
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

	err := s.Templates[tplName].ExecuteTemplate(w, "base", tpl)
	if err != nil {
		fmt.Println("errServe: ", err)
	}
}

func (a alertLevel) String() string {
	switch a {
	case alertLevelSuccess:
		return "success"
	case alertLevelDanger:
		return "danger"
	case alertLevelWarning:
		return "warning"
	}
	// maybe change later
	return "danger"
}

const (
	templatesBasePath = "templates/"
	templatesExt      = ".gohtml"
)

func (s *Server) parseTemplate(name, path string) {
	if path == "" {
		path = name
	}

	// if _, ok := s.Templates[name]; ok {
	// 	log.Panic("template", name)
	// 	return
	// }

	s.Templates[name] = template.Must(template.ParseFiles(templatesBasePath+"base"+templatesExt, templatesBasePath+name+templatesExt))
}
