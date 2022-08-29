package database

import (
	"database/sql"
	"errors"
	"log"
	"mealplanner/models"

	"github.com/lib/pq"
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

func (db DBConnection) SelectHousehold(HouseholdId string) (models.HouseHoldDB, error) {
	var householdmodel = models.HouseHoldDB{}
	query := `SELECT * FROM household WHERE householdid = $1`
	res := db.Con.QueryRow(query, HouseholdId).Scan(&householdmodel.Id, &householdmodel.HouseholdId, &householdmodel.GroceryListId, pq.Array(&householdmodel.Plans), pq.Array(&householdmodel.Meals))
	if res == sql.ErrNoRows {
		return householdmodel, errors.New("no id found")
	}

	return householdmodel, nil
}

func (db DBConnection) UpdateHousehold(household models.HouseHold) error {

	query := `UPDATE household SET grocerylist = $1 WHERE householdid = $2`
	_, err := db.Con.Exec(query, household.GroceryListId, household.HouseholdId)
	if err != nil {
		log.Println("Failed updating household")
		return err
	}
	return nil
}

func (db DBConnection) UpdateHouseholdArrays(HouseholdId string, updateField string, ids []int64) error {
	var values []int64
	if updateField == "meals" {
		values = ids
	} else if updateField == "plans" {
		values = ids
	}

	query := `UPDATE household SET ` + updateField + ` = $1 WHERE householdid = $2`
	_, err := db.Con.Exec(query, pq.Array(values), HouseholdId)
	if err != nil {
		log.Println("Failed updating household")
		return err
	}
	return nil
}

// func (db DBConnection) DeleteHouseHold(household models.HouseHold) error {
// 	query := `DELETE FROM household WHERE householdid = $1`
// }
