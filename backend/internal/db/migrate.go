package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB) {
	content, err := os.ReadFile("migrations/001_init.sql")
	if err != nil {
		log.Fatal("❌ Cannot read migration:", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Migrations applied")
}
