package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitMigration(t *testing.T) {
	defer clearTest()

	migrationSource := "file://migration"
	dbSource := "sqlite3://./zsync.db"

	// setup a sqlite database and run migrations
	InitMigration(migrationSource, dbSource)

	// check if the database is created
	f, err := os.Stat("./zsync.db")
	require.NoError(t, err)
	require.Equal(t, "zsync.db", f.Name())
}

func clearTest() {
	os.Remove("./zsync.db")
}
