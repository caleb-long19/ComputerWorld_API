package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UID       string         `gorm:"primaryKey" json:"uid"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Name      string         `gorm:"not null" json:"name"`
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Will implement two-factor authentication at a later date
}

func (User) TableName() string {
	return "user"
}

func (t *User) BeforeCreate(tx *gorm.DB) error {
	// If the UID is already set then just return.
	if t.UID != "" {
		return nil
	}
	newUuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	t.UID = newUuid.String()
	return nil
}
