package db

import (
	"gorm.io/gorm"
)

func init() {
	if globalDB == nil {
		globalDB = InitDB()
	}
}

var (
	globalDB *gorm.DB
)

// CreateConfig creates a new Config record
func CreateConfig(config Config) error {
	return globalDB.Create(&config).Error
}

// GetConfig retrieves a Config record
func GetConfig() (Config, error) {
	var config Config
	err := globalDB.First(&config).Error
	return config, err
}

// UpdateConfig updates an existing Config record
func UpdateConfig(config Config) error {
	return globalDB.Save(&config).Error
}

// DeleteConfig deletes a Config record
func DeleteConfig(config Config) error {
	return globalDB.Delete(&config).Error
}

// CreateURL creates a new URL record
func CreateURL(url URL) error {
	return globalDB.Create(&url).Error
}

// GetURL retrieves a URL record by key
func GetURL(key int) (URL, error) {
	var url URL
	err := globalDB.First(&url, key).Error
	return url, err
}

// UpdateURL updates an existing URL record
func UpdateURL(url URL) error {
	return globalDB.Save(&url).Error
}

// DeleteURL deletes a URL record
func DeleteURL(url URL) error {
	return globalDB.Delete(&url).Error
}

// CreateCommand creates a new Command record
func CreateCommand(command Command) error {
	return globalDB.Create(&command).Error
}

// GetCommand retrieves a Command record by key
func GetCommand(key int) (Command, error) {
	var command Command
	err := globalDB.First(&command, key).Error
	return command, err
}

func GetAllCommands() ([]Command, error) {
	var commands []Command
	err := globalDB.Find(&commands).Error
	return commands, err
}

// UpdateCommand updates an existing Command record
func UpdateCommand(command Command) error {
	return globalDB.Save(&command).Error
}

// DeleteCommand deletes a Command record
func DeleteCommand(command Command) error {
	return globalDB.Delete(&command).Error
}
