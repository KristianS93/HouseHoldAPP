package database

import (
	"database/sql"
	"errors"
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

func (db DBConnection) SelectGroceryList(household models.HouseHold) (models.HouseHold, error) {
	var householdmodel = models.HouseHold{}
	query := `SELECT * FROM household WHERE householdid = $1`
	res := db.Con.QueryRow(query, household.HouseholdId).Scan(&householdmodel.Id, &householdmodel.HouseholdId, &householdmodel.Meals, &householdmodel.GroceryListId)
	if res == sql.ErrNoRows {
		return householdmodel, errors.New("no id found")
	}

	return householdmodel, nil
}

func (db DBConnection) UpdateHousehold(household models.HouseHold, updateField string, itemString string) error {

	query := `UPDATE household SET ` + updateField + ` = $1 WHERE householdid = $2`
	var value string
	if updateField == "grocerylist" {
		value = household.GroceryListId
	} else if updateField == "meals" {
		value = itemString
	}
	_, err := db.Con.Exec(query, value, household.HouseholdId)
	if err != nil {
		log.Printf("Failed updating household with %s ", updateField)
		return err
	}
	return nil
}
