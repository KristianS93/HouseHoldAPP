package database

import (
	"errors"
	"log"
	"mealplanner/models"
	"strconv"

	"github.com/lib/pq"
)

func (db DBConnection) InsertMeal(meal *models.Meal) error {
	//Handle the array of items:
	var itemStr string
	for _, v := range meal.Items {
		x := strconv.Itoa(v.Id)
		itemStr = itemStr + x + ","
	}
	// fmt.Printf("Meal name: %s, desc: %s, items: %s", meal.MealName, meal.Description, itemStr)
	var id int
	query := `INSERT INTO meal (name, description, items) VALUES ($1, $2, $3) RETURNING id`
	err := db.Con.QueryRow(query, meal.MealName, meal.Description, itemStr).Scan(&id)
	if err != nil {
		return err
	}
	meal.Id = id
	return nil
}

func (db DBConnection) SelectMealId(id int) (models.Meal, string, error) {
	var meal = models.Meal{}
	var itemString string
	query := `SELECT * FROM meal WHERE id = $1`
	err := db.Con.QueryRow(query, id).Scan(&meal.Id, &meal.MealName, &meal.Description, &itemString)
	if err != nil {
		return meal, "", errors.New("no id found")
	}
	return meal, itemString, nil
}

func (db DBConnection) SelectMultipleItems(id []string) ([]models.Item, error) {
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
