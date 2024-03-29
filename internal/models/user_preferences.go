package models

type UserPreferences struct {
	BaseModel
	UserID    uint `json:"-"`
	User      User `json:"-"`
	VaultView int  `gorm:"not null; default: 0; check:vault_view = 0 or vault_view = 1" json:"vaultView"`
}
