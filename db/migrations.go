package db

import (
	"time"
)

// Config represents the config table
type Config struct {
	CronSeconds int `gorm:"not null"`
}

// URL represents the urls table
type URL struct {
	Key       int       `gorm:"primaryKey;autoIncrement"`
	URL       string    `gorm:"not null"`
	Platform  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

// Command represents the commands table
type Command struct {
	Key          int       `gorm:"primaryKey;autoIncrement"`
	Command      string    `gorm:"not null"`
	Frequency    int       `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	LastCalledAt time.Time `gorm:"not null"`
}
