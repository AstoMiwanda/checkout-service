package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Stock struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	ProductID uuid.UUID `gorm:"references:products"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	Qty       int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt sql.NullTime
	DeletedAt gorm.DeletedAt
}
