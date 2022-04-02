package models

import (
	"gorm.io/gorm"
)

type Vault struct {
	gorm.Model
	UserID uint `gorm: "not null; uniqueIndex"`
	User   User
}
