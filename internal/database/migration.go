package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/kerupuksambel/portfolio-be/utils"
)

func Migrate() {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Printf("[WARN] Migration error due to database error. %s", err)
		return
	}
	defer db.Close()

	migrationDir := fmt.Sprintf("file://%s/internal/database/migrations", utils.ExecPath())
	m, err := migrate.New(migrationDir, url)
	if err != nil {
		log.Printf("[WARN] Migration error due to migration initialization error. %s", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration error: %s", err)
	}
}
