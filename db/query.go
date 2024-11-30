package db

import (
	"gorm.io/gorm"
)

// CreateConfig creates a new Config record
func CreateConfig(db *gorm.DB, config Config) error {
	return db.Create(&config).Error
}

// GetConfig retrieves a Config record
func GetConfig(db *gorm.DB) (Config, error) {
	var config Config
	err := db.First(&config).Error
	return config, err
}

// UpdateConfig updates an existing Config record
func UpdateConfig(db *gorm.DB, config Config) error {
	return db.Save(&config).Error
}

// DeleteConfig deletes a Config record
func DeleteConfig(db *gorm.DB, config Config) error {
	return db.Delete(&config).Error
}

// CreateURL creates a new URL record
func CreateURL(db *gorm.DB, url URL) error {
	return db.Create(&url).Error
}

// GetURL retrieves a URL record by key
func GetURL(db *gorm.DB, key int) (URL, error) {
	var url URL
	err := db.First(&url, key).Error
	return url, err
}

// UpdateURL updates an existing URL record
func UpdateURL(db *gorm.DB, url URL) error {
	return db.Save(&url).Error
}

// DeleteURL deletes a URL record
func DeleteURL(db *gorm.DB, url URL) error {
	return db.Delete(&url).Error
}

// CreateCommand creates a new Command record
func CreateCommand(db *gorm.DB, command Command) error {
	return db.Create(&command).Error
}

// GetCommand retrieves a Command record by key
func GetCommand(db *gorm.DB, key int) (Command, error) {
	var command Command
	err := db.First(&command, key).Error
	return command, err
}

// UpdateCommand updates an existing Command record
func UpdateCommand(db *gorm.DB, command Command) error {
	return db.Save(&command).Error
}

// DeleteCommand deletes a Command record
func DeleteCommand(db *gorm.DB, command Command) error {
	return db.Delete(&command).Error
}
