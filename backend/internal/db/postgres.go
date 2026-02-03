package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var db *sqlx.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = sqlx.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			log.Println("✅ Connected to PostgreSQL")
			return db
		}

		log.Printf("⏳ Waiting for DB... attempt %d/10\n", i)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Could not connect to DB after multiple attempts")
	return nil
}
