package model

import (
	"database/sql"
	"io/ioutil"
	"log"

	// Pull in the sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// initDatabase opens the database handle
func initDatabase(dbPath, schemaPath string) *sql.DB {
	log.Printf("initializing database at %s with schema at %s", dbPath, schemaPath)
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	schema, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("error loading schema sql: %v", err)
	}

	_, err = database.Exec(string(schema))
	if err != nil {
		log.Fatalf("error applying schema sql: %v", err)
	}

	return database
}
