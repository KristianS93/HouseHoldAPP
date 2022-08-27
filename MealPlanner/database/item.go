package database

import (
	"fmt"
	"log"
)

func (db DBConnection) InsertItems(valueStr string, data []interface{}) error {
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
	}

	rows, err := tx.Query(query, data...)
	if err != nil {
		log.Println("Error in query: ", err)
		tx.Rollback()
	}

	for rows.Next() {
		var itemsId int
		rows.Scan(&itemsId)
		ids = append(ids, itemsId)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(ids)

	return nil
}
