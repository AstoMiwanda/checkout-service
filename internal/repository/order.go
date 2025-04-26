package repository

import (
	"checkout-service/pkg/model"
	"context"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(context.Context, model.Order) (model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r orderRepository) Create(ctx context.Context, order model.Order) (result model.Order, err error) {
	err = r.db.WithContext(ctx).Create(&order).Scan(&result).Error
	if err != nil {
		return order, err
	}
	return result, nil
}
