package web

import (
	"net/http"
)

type alertLevel uint8

const (
	Success alertLevel = iota
	Warning
	Danger
)

type Alert struct {
	Level   string
	Message string
}

func (a alertLevel) String() string {
	switch a {
	case Success:
		return "success"
	case Danger:
		return "danger"
	case Warning:
		return "warning"
	default:
		return ""
	}
}

// maybe refactor constants for error messages on alerts
func addAlert(w http.ResponseWriter, r *http.Request, alertType alertLevel, message string) {
	c := http.Cookie{
		Name:   alertType.String(),
		Value:  message,
		MaxAge: 3,
	}
	http.SetCookie(w, &c)
}

func getAlert(w http.ResponseWriter, r *http.Request) []Alert {
	var Alerts []Alert
	if cWarning, err := r.Cookie(Warning.String()); err == nil {
		Alerts = append(Alerts, Alert{
			Level:   cWarning.Name,
			Message: cWarning.Value,
		})
	}
	if cDanger, err := r.Cookie(Danger.String()); err == nil {
		Alerts = append(Alerts, Alert{
			Level:   cDanger.Name,
			Message: cDanger.Value,
		})
	}
	if cSuccess, err := r.Cookie(Success.String()); err == nil {
		Alerts = append(Alerts, Alert{
			Level:   cSuccess.Name,
			Message: cSuccess.Value,
		})
	}

	return Alerts
}
