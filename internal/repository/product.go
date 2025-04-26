package repository

import (
	"checkout-service/pkg/model"
	"checkout-service/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(context.Context, model.Product) (model.Product, error)
	Update(context.Context, model.Product) (model.Product, error)
	GetDetail(context.Context, uuid.UUID) (model.Product, error)
	GetList(context.Context, utils.Pagination, map[string]interface{}) ([]model.Product, int64, error)
	Delete(context.Context, uuid.UUID) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r productRepository) Create(ctx context.Context, product model.Product) (result model.Product, err error) {
	err = r.db.WithContext(ctx).Create(&model.Product{
		Sku:       product.Sku,
		Name:      product.Name,
		Price:     product.Price,
		IsActive:  product.IsActive,
		CreatedAt: time.Now(),
	}).Scan(&result).Error
	if err != nil {
		return product, err
	}
	return result, nil
}

func (r productRepository) Update(ctx context.Context, product model.Product) (result model.Product, err error) {
	updates := map[string]interface{}{
		"sku":        product.Sku,
		"name":       product.Name,
		"price":      product.Price,
		"is_active":  product.IsActive,
		"updated_at": sql.NullTime{Time: time.Now(), Valid: true},
	}

	err = r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ?", product.ID).
		Updates(updates).
		Scan(&result).Error
	if err != nil {
		return product, err
	}
	return result, nil
}

func (r productRepository) GetDetail(ctx context.Context, id uuid.UUID) (result model.Product, err error) {
	err = r.db.WithContext(ctx).
		Where(&model.Product{ID: id}).
		First(&result, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Product{}, fmt.Errorf("product with id %s not found", id)
		}
		return model.Product{}, fmt.Errorf("failed to get product detail: %w", err)
	}
	return result, nil
}

func (r productRepository) GetList(ctx context.Context, pagination utils.Pagination, where map[string]interface{}) (items []model.Product, count int64, err error) {
	query := r.db.WithContext(ctx).Model(&model.Product{})

	if len(where) > 0 {
		query = query.Where(where)
	}

	if pagination.Field == "" {
		pagination.Field = "created_at"
	}

	if pagination.Sort == "" {
		pagination.Sort = "ASC"
	}

	orderBy := fmt.Sprintf("%s %s", pagination.Field, pagination.Sort)
	offset := (pagination.Page - 1) * pagination.Limit
	limitBuilder := query.Limit(int(pagination.Limit)).Offset(int(offset)).Order(orderBy)

	result := limitBuilder.Find(&items)
	if result.Error != nil {
		return nil, count, result.Error
	}

	err = query.Count(&count).Error
	if err != nil {
		return nil, count, err
	}

	return items, count, err
}

func (r productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Product{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete product: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("product with id %s not found", id)
	}
	return nil
}
