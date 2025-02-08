package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

var GlobalStore *Store

type Store struct {
	Querier
	db *sql.DB
}

func NewStore(dbPath, migrationPath string) *Store {
	db, err := sql.Open("sqlite", fmt.Sprintf("./%s", dbPath))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationPath),
		fmt.Sprintf("sqlite3://%s", dbPath),
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	GlobalStore := &Store{
		db:      db,
		Querier: New(db),
	}

	return GlobalStore
}
