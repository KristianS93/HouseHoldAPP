package database

import (
	"errors"
	"log"
	"mealplanner/models"

	"github.com/lib/pq"
)

func (db DBConnection) InsertMeal(meal *models.Meal, itemIds []int64) error {
	var id int64
	query := `INSERT INTO meal (name, description, items) VALUES ($1, $2, $3) RETURNING id`
	err := db.Con.QueryRow(query, meal.MealName, meal.Description, pq.Array(itemIds)).Scan(&id)
	if err != nil {
		return err
	}
	meal.Id = id
	return nil
}

func (db DBConnection) SelectMealId(id int64) (models.MealDB, error) {
	var meal = models.MealDB{}
	query := `SELECT * FROM meal WHERE id = $1`
	err := db.Con.QueryRow(query, id).Scan(&meal.Id, &meal.MealName, &meal.Description, pq.Array(&meal.Items))
	if err != nil {
		log.Println(err)
		return meal, errors.New("no id found")
	}
	return meal, nil
}

func (db DBConnection) SelectMultipleItems(id []int64) ([]models.Item, error) {
	var items []models.Item

	buildQuery := "SELECT * FROM item WHERE id = ANY($1::int[])"

	rows, err := db.Con.Query(buildQuery, pq.Array(id))
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var item models.Item
		rows.Scan(&item.Id, &item.ItemName, &item.Quantity, &item.Unit)
		items = append(items, item)
	}
	return items, nil
}

func (db DBConnection) SelectMultipleMeals(id []int64) ([]models.MealDB, error) {
	var meals []models.MealDB

	buildQuery := `SELECT * FROM meal WHERE id = ANY($1::int[])`

	rows, err := db.Con.Query(buildQuery, pq.Array(id))
	if err != nil {
		return meals, err
	}

	for rows.Next() {
		var meal models.MealDB
		rows.Scan(&meal.Id, &meal.MealName, &meal.Description, pq.Array(&meal.Items))
		meals = append(meals, meal)
	}

	return meals, nil
}

func (db DBConnection) DeleteMeal(mealId int64) error {
	query := `DELETE FROM meal WHERE id = $1`
	_, err := db.Con.Exec(query, mealId)
	if err != nil {
		return err
	}
	return nil
}

func (db DBConnection) DeleteMeals(ids []int64) error {
	query := `DELETE FROM meal WHERE id = ANY($1::int[])`
	_, err := db.Con.Exec(query, pq.Array(ids))
	if err != nil {
		log.Println("error executing query")
		return err
	}
	return nil
}

func (db DBConnection) UpdateMeal(meal models.Meal) error {
	query := `UPDATE meal SET name = $1, description = $2 WHERE id = $3`
	_, err := db.Con.Exec(query, meal.MealName, meal.Description, meal.Id)
	if err != nil {
		return err
	}
	return err
}
