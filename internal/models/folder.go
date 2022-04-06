package models

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	VaultID     uint `gorm:"not null"`
	Vault       Vault
	FolderID    *uint
	Folder      *Folder
	Name        string `gorm:"not null"`
	Description string
}
