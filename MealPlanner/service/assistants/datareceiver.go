package assistants

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodeData(r *http.Request, datastructure any) error {
	err := json.NewDecoder(r.Body).Decode(&datastructure)
	if err != nil {
		log.Println("Error: Parsing Json")
		return err
	}
	return nil
}
