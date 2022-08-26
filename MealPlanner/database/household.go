package database

import (
	"log"
	"mealplanner/models"
)

func (db DBConnection) InsertHousehold(householdid models.HouseHold) error {
	//Insert into household:
	query := `INSERT INTO household (householdid) VALUES ($1)`
	_, err := db.Con.Exec(query, householdid.HouseholdId)
	if err != nil {
		log.Println("Error inserting household")
		return err
	}
	return nil
}
