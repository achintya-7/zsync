package db

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// InitDbAndMigration sets up the SQLite database and runs migrations
func InitDbAndMigration() {
	// setup a sqlite database using GORM
	db, err := gorm.Open(sqlite.Open("zsync.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// run migrations using GORM
	err = db.AutoMigrate(&Config{}, &URL{}, &Command{})
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() *gorm.DB {
	// setup a sqlite database using GORM
	db, err := gorm.Open(sqlite.Open("zsync.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
