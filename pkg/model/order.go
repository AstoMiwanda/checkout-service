package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CustomerID    uuid.UUID `gorm:"references:customers"`
	Customer      Customer  `gorm:"foreignKey:CustomerID"`
	Status        string    `gorm:"type:varchar(20);default:'pending'"`
	TotalAmount   float64
	TotalDiscount float64
	TotalPayment  float64
	OrderDiscount []OrderDiscount
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     sql.NullTime
	DeletedAt     gorm.DeletedAt
	OrderDetails  []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	OrderID   uuid.UUID `gorm:"references:orders"`
	Order     Order     `gorm:"foreignKey:OrderID"`
	ProductID uuid.UUID `gorm:"references:products"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	Subtotal  float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt sql.NullTime
	DeletedAt gorm.DeletedAt
}

type OrderDiscount struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	OrderID   uuid.UUID `gorm:"not null"`
	ProductId uuid.UUID `gorm:"not null"`
	Name      string
	Qty       int
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt sql.NullTime
	DeletedAt gorm.DeletedAt
}
