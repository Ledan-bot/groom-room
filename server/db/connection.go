package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sqlConnect := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", sqlConnect)
	if err != nil {
		log.Fatal("Error Connecting to DB", err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	return db
}
