package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Key      string `gorm:"not null"`
	Username string `gorm:"not null; uniqueIndex"`
	Password string `gorm:"not null"`
	Email    string
}
