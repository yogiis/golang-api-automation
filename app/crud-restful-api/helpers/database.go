package helpers

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	passworddb := os.Getenv("PASSWORDDB")
	db, err := sql.Open("mysql", "root:"+passworddb+"@tcp(localhost:3306)/crudapi")
	LogPanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
