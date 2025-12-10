package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DBCreater struct{}

func NewDBCreater() *DBCreater {
	return &DBCreater{}
}

func (d *DBCreater) InitDatabaseIfNeeded() {
	db, err := sql.Open("postgres", os.Getenv("DB_URL_ADMIN"))
	if err != nil {
		log.Fatal("Error connecting to Postgres:", err)
	}
	defer db.Close()

	dbname := os.Getenv("DB_NAME")
	testDbName := dbname + "_test"
	d.createDB(db, dbname)
	d.createDB(db, testDbName)
}

func (d *DBCreater) createDB(db *sql.DB, dbname string) {
	exists := d.dbExists(db, dbname)

	if !exists {
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			log.Fatal("Error creating database:", err)
		}
		log.Printf("Database '%s' created successfully!", dbname)
	}
}

func (d *DBCreater) dbExists(db *sql.DB, dbname string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT datname FROM pg_database WHERE datname = $1)"
	err := db.QueryRow(query, dbname).Scan(&exists)
	if err != nil {
		log.Fatalf("Error verifying database existence: %v", err)
	}
	return exists
}
