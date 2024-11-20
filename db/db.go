package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

// setup a sqlite database and run migrations
func InitMigration(migrationSource string, dbSource string) {
	// setup a sqlite database
	db, err := sql.Open("sqlite3", "./zsync.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// run migrations
	err = runDbMigrations(migrationSource, dbSource)
	if err != nil {
		log.Fatal(err)
	}
}

func runDbMigrations(migrationUrl string, dbSource string) error {
	// run migrations
	migration, err := migrate.New(migrationUrl, dbSource)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil {
		return err
	}

	return nil
}
