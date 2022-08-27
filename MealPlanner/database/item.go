package database

import (
	"errors"
	"fmt"
	"log"
	"mealplanner/models"

	"github.com/lib/pq"
)

func (db DBConnection) InsertItems(valueStr string, data []interface{}) ([]int, error) {
	var ids []int
	query := fmt.Sprintf("INSERT INTO item (name, quantity, unit) VALUES %s RETURNING id", valueStr)
	// VIRKER
	// rows, err := db.Con.Query(query, data...)
	// if err != nil {
	// 	log.Println("Error in query: ", err)
	// }

	// for rows.Next() {
	// 	var itemsId int
	// 	rows.Scan(&itemsId)
	// 	ids = append(ids, itemsId)
	// }

	tx, err := db.Con.Begin()
	if err != nil {
		log.Println(err)
		return ids, err
	}

	rows, err := tx.Query(query, data...)
	if err != nil {
		log.Println("Error in query: ", err)
		tx.Rollback()
		return ids, err
	}

	for rows.Next() {
		var itemsId int
		rows.Scan(&itemsId)
		ids = append(ids, itemsId)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return ids, err
	}

	return ids, nil
}

func (db DBConnection) DeleteItems(ids []int) error {
	items, err := db.SelectMultipleItems(ids)
	if err != nil {
		log.Println("Error checking if items exist")
		return err
	}
	if len(items) == 0 {
		return errors.New("items does not exist")
	}
	query := `DELETE FROM item WHERE id = ANY($1::int[])`
	_, err = db.Con.Exec(query, pq.Array(ids))
	if err != nil {
		log.Println("error executing query items")
		return err
	}
	return nil
}

func (db DBConnection) UpdateItem(data models.Item) error {
	query := `UPDATE item SET name = $1, quantity = $2, unit = $3  WHERE id = $4`
	_, err := db.Con.Exec(query, data.ItemName, data.Quantity, data.Unit, data.Id)
	if err != nil {
		log.Println("Failed updating item")
		return err
	}
	return nil
}
