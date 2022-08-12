package web

import (
	"net/http"
)

type alertLevel uint8

const (
	alertLevelSuccess alertLevel = iota
	alertLevelWarning
	alertLevelDanger
)

type Alert struct {
	Level   string
	Message string
}

func (a alertLevel) String() string {
	switch a {
	case alertLevelSuccess:
		return "success"
	case alertLevelDanger:
		return "danger"
	case alertLevelWarning:
		return "warning"
	default:
		return ""
	}
}

func addAlert(w http.ResponseWriter, r *http.Request, alertType alertLevel, message string) {
	c := http.Cookie{
		Name:   alertType.String(),
		Value:  message,
		MaxAge: 5,
	}
	http.SetCookie(w, &c)
}

func getAlert(r *http.Request) []Alert {
	var Alerts []Alert
	if cWarning, err := r.Cookie(alertLevelWarning.String()); err == nil {
		Alerts = append(Alerts, Alert{
			Level:   cWarning.Name,
			Message: cWarning.Value,
		})
	}
	if cDanger, err := r.Cookie(alertLevelDanger.String()); err == nil {
		Alerts = append(Alerts, Alert{
			Level:   cDanger.Name,
			Message: cDanger.Value,
		})
	}
	if cSuccess, err := r.Cookie(alertLevelSuccess.String()); err == nil {
		Alerts = append(Alerts, Alert{
			Level:   cSuccess.Name,
			Message: cSuccess.Value,
		})
	}

	return Alerts
}
