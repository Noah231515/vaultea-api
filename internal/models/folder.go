package models

type Folder struct {
	BaseModel
	VaultID     uint    `gorm:"not null" json:"-"`
	Vault       Vault   `json:"-"`
	FolderID    *uint   `json:"-"`
	Folder      *Folder `json:"-"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Starred     bool    `json:"starred"`
}
