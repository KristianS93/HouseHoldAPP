package database

import (
	"errors"
	"log"
	"mealplanner/models"

	"github.com/lib/pq"
)

func (db DBConnection) InsertMeal(meal *models.Meal, itemIds []int) error {
	var id int
	query := `INSERT INTO meal (name, description, items) VALUES ($1, $2, $3) RETURNING id`
	err := db.Con.QueryRow(query, meal.MealName, meal.Description, pq.Array(itemIds)).Scan(&id)
	if err != nil {
		return err
	}
	meal.Id = id
	return nil
}

func (db DBConnection) SelectMealId(id int) (models.MealDB, error) {
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

func (db DBConnection) DeleteMeal(mealId int) error {
	query := `DELETE FROM meal WHERE id = $1`
	_, err := db.Con.Exec(query, mealId)
	if err != nil {
		return err
	}
	return nil
}
