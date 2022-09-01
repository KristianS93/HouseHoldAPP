package database

import (
	"errors"
	"log"
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

func (db DBConnection) DeletePlans(ids []int64) error {
	query := `DELETE FROM plan WHERE id = ANY($1::int[])`
	_, err := db.Con.Exec(query, pq.Array(ids))
	if err != nil {
		log.Println("error executing query")
		return err
	}
	return nil
}

func (db DBConnection) UpdatePlan(planData models.PlanDB) error {
	query := `UPDATE plan SET weekno = $1, meals = $2 WHERE id = $3`
	_, err := db.Con.Exec(query, planData.WeekNo, pq.Array(planData.Meals), planData.Id)
	if err != nil {
		return err
	}
	return nil
}

func (db DBConnection) SelectPlan(weekno int64, household string) (models.PlanDB, error) {
	var rPlanData models.PlanDB
	query := `SELECT * FROM plan WHERE weekno = $1 AND householdid = $2`
	err := db.Con.QueryRow(query, weekno, household).Scan(&rPlanData.Id, &rPlanData.WeekNo, &rPlanData.HouseHoldId, pq.Array(&rPlanData.Meals))
	if err != nil {
		log.Println(err)
		return rPlanData, errors.New("error selecting plan")
	}
	return rPlanData, nil
}

func (db DBConnection) TestMultipleWeekno(weekno int, household string) bool {
	var count int
	query := `SELECT COUNT(*) FROM plan WHERE weekno = $1 AND householdid = $2`
	err := db.Con.QueryRow(query, weekno, household).Scan(&count)
	if err != nil {
		log.Println(err)
		return false
	}
	if count > 0 {
		return false
	}
	return true
}
