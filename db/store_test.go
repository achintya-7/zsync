package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDbInit(t *testing.T) {
	// initialize the store
	store := NewStore("zsync.db", "migrations")

	defer os.Remove("./zsync.db")
	defer store.db.Close()

	// insert a url
	ctx := context.Background()

	insertUrlParams := InsertUrlParams{
		Url:      "https://www.google.com",
		Platform: "mac",
	}

	err := store.InsertUrl(ctx, insertUrlParams)
	require.NoError(t, err)

	// get all the urls
	urls, err := store.GetAllUrls(ctx)
	require.NoError(t, err)
	require.Greater(t, len(urls), 1)
}
