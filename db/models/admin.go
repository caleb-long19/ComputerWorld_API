package models

type Admin struct {
	AdminID  int    `gorm:"autoIncrement;primaryKey" json:"admin_id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
}
