package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading ENV File")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Error Connecting to DB", err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
