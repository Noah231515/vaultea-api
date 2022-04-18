package models

type User struct {
	BaseModel
	Key      string `gorm:"not null"`
	Username string `gorm:"not null; uniqueIndex"`
	Password string `gorm:"not null"`
	Email    string
}
