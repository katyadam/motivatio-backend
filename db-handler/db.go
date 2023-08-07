package db_handler

import (
	"database/sql"
	"log"
)

func getDb() *sql.DB {
	db, err := sql.Open(
		"postgre",
		"username:password@tcp(127.0.0.1:3306)/database_name",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	return db
}
