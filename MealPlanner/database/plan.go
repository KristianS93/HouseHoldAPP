package database

import (
	"mealplanner/models"

	"github.com/lib/pq"
)

func (db DBConnection) CreatePlan(data *models.Plan, mealIds []int64) error {
	var id int64
	query := `INSERT INTO plan (weekno, householdid, meals) VALUES ($1, $2, $3) RETURNING id`
	err := db.Con.QueryRow(query, data.WeekNo, data.HouseHoldId, pq.Array(mealIds)).Scan(&id)
	if err != nil {
		return err
	}
	data.Id = id
	return nil
}

func (db DBConnection) DeletePlan(planId int64) error {
	query := `DELETE FROM plan WHERE id = $1`
	_, err := db.Con.Exec(query, planId)
	if err != nil {
		return err
	}
	return nil
}
