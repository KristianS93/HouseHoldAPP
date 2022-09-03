package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// DecodeRequest decodes the json body of a request
// into the provided destination, which can be any type
// if an error occurs during decoding, that error is returned.
//
// The destination MUST be sent as a pointer, ie. "&data"
// otherwise the function will not result in anything.
func DecodeRequest(r *http.Request, dest any) error {
	err := json.NewDecoder(r.Body).Decode(&dest)
	if err != nil {
		log.Println("Failed to decode request body: ", err)
		return err
	}
	return nil
}

func CheckResult(result sql.Result) error {
	ra, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if ra != 1 {
		return fmt.Errorf("%d rows were affected by the query", ra)
	}
	return nil
}
