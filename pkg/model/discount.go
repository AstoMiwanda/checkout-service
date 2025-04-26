package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Discount struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	Type          string    `gorm:"not null"`
	Description   string
	DiscountValue float64
	IsActive      bool `gorm:"default:true"`
	ValidFrom     time.Time
	ValidTo       time.Time
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     sql.NullTime
	DeletedAt     gorm.DeletedAt
	DiscountRules []DiscountRule `gorm:"foreignKey:DiscountID"`
}

type DiscountRule struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	DiscountID       uuid.UUID `gorm:"references:discounts"`
	Discount         Discount  `gorm:"foreignKey:DiscountID"`
	ProductID        uuid.UUID `gorm:"references:products"`
	Product          Product   `gorm:"foreignKey:ProductID"`
	Role             string    `gorm:"type:varchar(20);not null"`
	Quantity         int       `gorm:"not null"`
	QuantityOperator string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        sql.NullTime
	DeletedAt        gorm.DeletedAt
}
