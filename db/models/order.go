package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	UID         string         `gorm:"primaryKey" json:"uid"`
	OrderRef    string         `json:"order_ref"`
	OrderAmount int            `json:"order_amount"`
	ProductUID  string         `json:"product_uid"`
	OrderPrice  float64        `json:"order_price"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`

	Product Product `gorm:"foreignKey:ProductID"` // Relationship to Product
}

func (Order) TableName() string {
	return "orders"
}

func (t *Order) BeforeCreate(tx *gorm.DB) error {
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
