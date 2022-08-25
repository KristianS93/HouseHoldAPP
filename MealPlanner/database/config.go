package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const (
//
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "postgres"
//	dbname   = "postgres"
//
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "MpAdmin"
// 	password = "MpPass"
// 	dbname   = "Mp"
// )

func Connect() {
	err := godotenv.Load("./database/.env")
	if err != nil {
		log.Fatal("Error loading env.")
	}
	usr := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("MP_HOST")
	port := os.Getenv("MP_PORT")

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)'
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, usr, pw, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
