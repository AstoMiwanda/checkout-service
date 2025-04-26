package repository

import (
	"checkout-service/pkg/model"
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscountRepository interface {
	GetDiscountRuleByProduct(ctx context.Context, productId uuid.UUID) ([]model.DiscountRule, error)
	GetDiscountById(ctx context.Context, discountId uuid.UUID) (model.Discount, error)
}

type discountRepository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *discountRepository {
	return &discountRepository{
		db: db,
	}
}

func (r discountRepository) GetDiscountRuleByProduct(ctx context.Context, productId uuid.UUID) (result []model.DiscountRule, err error) {
	var discountRule model.DiscountRule
	err = r.db.WithContext(ctx).
		Model(&model.DiscountRule{}).
		Where(&model.DiscountRule{ProductID: productId}).
		Scan(&discountRule).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get discount detail: %w", err)
	}

	err = r.db.WithContext(ctx).
		Model(&model.DiscountRule{}).
		Where(&model.DiscountRule{DiscountID: discountRule.DiscountID}).
		Find(&result).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get discount detail: %w", err)
	}
	return result, nil
}

func (r discountRepository) GetDiscountById(ctx context.Context, discountId uuid.UUID) (result model.Discount, err error) {
	err = r.db.WithContext(ctx).
		Model(&model.Discount{}).
		Where(&model.Discount{ID: discountId}).
		Find(&result).Error
	if err != nil {
		return model.Discount{}, fmt.Errorf("failed to get discount detail: %w", err)
	}
	return result, nil
}
