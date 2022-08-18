package web

import (
	"log"
	"net/http"
)

type alertLevel uint8

const (
	Success alertLevel = iota
	Warning
	Danger
)

const (
	InternalError = "Internal error."
	EmptyField    = "One field was empty, please fill all fields appropriately."
	ItemAdded     = "Item added successfully."
	ListCleared   = "List cleared successfully."
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
func addAlert(w http.ResponseWriter, alertType alertLevel, alertMessage string) {
	c := http.Cookie{
		Name:   alertType.String(),
		Value:  alertMessage,
		MaxAge: 3,
	}
	http.SetCookie(w, &c)
}

func GetAlert(w http.ResponseWriter, r *http.Request) []Alert {
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

// AlertLog adds an alert and logs a message, the log message should contain
// any relevant errors for the given function call.
// alertType must be one of the predefined constants of type alertLevel
// alertMessage should utilize predefined constants such as InternalError,
// however, it can also take string literals.
func AlertLog(w http.ResponseWriter, alertType alertLevel, alertMessage, logMessage string) {
	log.Println(logMessage)
	addAlert(w, alertType, alertMessage)
}
