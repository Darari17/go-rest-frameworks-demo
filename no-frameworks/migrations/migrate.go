package migrations

import (
	"database/sql"
	"log"
	"os"
)

func Migrate(db *sql.DB) error {
	query, err := os.ReadFile("migrations/init.sql")
	if err != nil {
		log.Fatal("Failed to read migration file:", err)
	}

	if _, err := db.Exec(string(query)); err != nil {
		log.Fatal("Migration failed:", err)
	}

	return nil
}
