package models

type User struct {
	UserID   int    `gorm:"autoIncrement;primaryKey" json:"user_id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`

	// Will implement two-factor authentication at a later date
}
