package meal

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// const host = "localhost"
// const port = 27021

type Database struct {
	Conn *sql.DB
}

func DbConnection(username, password, database string) (Database, error) {
	db := Database{}

	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	conn, err := sql.Open("postgres", "PG_DATABASE://PG_USER:PG_PASS@localhost:27021/postgres")
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
