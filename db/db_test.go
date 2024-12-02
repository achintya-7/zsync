package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitMigration(t *testing.T) {
	defer clearTest()

	InitDbAndMigration()

	// check if the database is created
	f, err := os.Stat("./zsync.db")
	require.NoError(t, err)
	require.Equal(t, "zsync.db", f.Name())
}

func clearTest() {
	os.Remove("./zsync.db")
}
