package models

import (
	"time"
)

type Password struct {
	BaseModel
	VaultID     uint  `gorm:"not null" json:"-"`
	Vault       Vault `json:"-"`
	FolderID    uint  `json:"folderId`
	Folder      Folder
	Username    string    `gorm:"not null" json:"username"`
	Password    string    `gorm:"not null" json:"password"`
	Description string    `json:"description"`
	ExpireDate  time.Time `json:"expireDate"`
}
