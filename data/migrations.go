package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func RunMigrations(db *sql.DB, migrationsPath string) error {
	files, err := filepath.Glob(filepath.Join(migrationsPath, "*.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migrations: %w", err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}
		log.Printf("Executed migration: %s", file)
	}

	return nil
}
