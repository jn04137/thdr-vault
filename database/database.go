package database

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dbPath string) (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Couldn't create/connect to sqlite.db: %v", err.Error())
		return nil, err
	}

	initializeTables(dbConn)
	return dbConn, err
}

func initializeTables(dbConn *sql.DB) {
	createApplicationTable := `CREATE TABLE IF NOT EXISTS application (
		id INTEGER PRIMARY KEY NOT NULL,
		app_name VARCHAR(120) NOT NULL
	);`
	_, err := dbConn.Exec(createApplicationTable)
	if err != nil {
		log.Fatalf("Not able to create application table: %v", err.Error())
	}

	createSecretsTable := `CREATE TABLE IF NOT EXISTS secret (
		id INTEGER PRIMARY KEY NOT NULL,
		key VARCHAR(100) NOT NULL UNIQUE,
		value VARCHAR(250) NOT NULL,
		check(key <> ''),
		check(value <> '')
	);`

	_, err = dbConn.Exec(createSecretsTable)
	if err != nil {
		log.Fatalf("Failed to create secret table: %v", err.Error())
	}
}

