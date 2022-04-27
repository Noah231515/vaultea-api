package models

type Vault struct {
	BaseModel
	UserID uint `gorm: "not null; uniqueIndex"`
	User   User
}
