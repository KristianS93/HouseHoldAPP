package database

import (
	"fmt"
	"mealplanner/models"
	"strconv"
)

func (db DBConnection) InsertMeal(meal *models.Meal) error {
	//Handle the array of items:
	var itemStr string
	for _, v := range meal.Items {
		x := strconv.Itoa(v.Id)
		itemStr = itemStr + x + ","
	}
	fmt.Printf("Meal name: %s, desc: %s, items: %s", meal.MealName, meal.Description, itemStr)
	var id int
	query := `INSERT INTO meal (name, description, items) VALUES ($1, $2, $3) RETURNING id`
	err := db.Con.QueryRow(query, meal.MealName, meal.Description, itemStr).Scan(&id)
	if err != nil {
		return err
	}
	meal.Id = id
	return nil
}
