package models

import (
	"time"

	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	VaultID     uint
	Vault       Vault `gorm:"not null"`
	FolderID    uint
	Folder      Folder
	Username    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Description string
	ExpireDate  time.Time
}
