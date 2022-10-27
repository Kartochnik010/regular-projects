package internal

import (
	"database/sql"
	"log"
)

func DB_conn() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=0000 dbname=productdb sslmode=disable")
	if err != nil {
		log.Fatal(err, " could not open database")
	}
	return db
}
