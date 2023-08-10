package dbHandler

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// GetDb
/**
Factory for db connection
*/
func GetDb() *sql.DB {
	db, err := sql.Open(
		"postgres",
		"user=postgres password=Motivatioo dbname=postgres host=127.0.0.1 port=5432 sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
