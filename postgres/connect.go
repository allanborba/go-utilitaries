package postgres

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func ConnectPostgres() *sql.DB {
	return connectToDB(os.Getenv("DB_URL"))
}

func connectToDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Printf("database connection pool established")

	return db
}
