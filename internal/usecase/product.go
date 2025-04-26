package usecase

import (
	"checkout-service/internal/repository"
	"checkout-service/internal/rest/request"
	"checkout-service/internal/rest/response"
	"checkout-service/pkg/model"
	"checkout-service/pkg/utils"
	"context"
	"github.com/google/uuid"
	"math"
)

type IProductService interface {
	CreateProduct(context.Context, request.CreateProductRequest) (*response.GetProductResponse, *response.ErrorResponse)
	GetDetailProduct(context.Context, string) (*response.GetProductResponse, *response.ErrorResponse)
	GetListProduct(context.Context, request.GetProductListRequest) (*response.GetProductListResponse, *response.ErrorResponse)
	UpdateProduct(context.Context, request.UpdateProductRequest) (*response.GetProductResponse, *response.ErrorResponse)
	DeleteProduct(context.Context, string) *response.ErrorResponse
}

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req request.CreateProductRequest) (*response.GetProductResponse, *response.ErrorResponse) {
	payload := model.Product{
		Sku:      req.Sku,
		Name:     req.Name,
		Price:    req.Price,
		IsActive: req.IsActive,
	}
	product, err := s.ProductRepository.Create(ctx, payload)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	result := ProductMapping(product)

	return &response.GetProductResponse{
		Product: result,
	}, nil
}

func (s *ProductService) GetDetailProduct(ctx context.Context, id string) (*response.GetProductResponse, *response.ErrorResponse) {
	idProduct, err := uuid.Parse(id)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	product, err := s.ProductRepository.GetDetail(ctx, idProduct)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	result := ProductMapping(product)

	return &response.GetProductResponse{
		Product: result,
	}, nil
}

func (s *ProductService) GetListProduct(ctx context.Context, req request.GetProductListRequest) (*response.GetProductListResponse, *response.ErrorResponse) {
	pagination := utils.GeneratePaginationFromRequest(req.Limit, req.Page, req.Field, req.Sort)
	products, count, err := s.ProductRepository.GetList(ctx, pagination, req.Where)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	var productList []response.Product
	for _, product := range products {
		productList = append(productList, ProductMapping(product))
	}

	return &response.GetProductListResponse{
		Page:      req.Page,
		Limit:     req.Limit,
		Total:     count,
		TotalPage: int64(math.Ceil(float64(count) / float64(pagination.Limit))),
		Products:  productList,
	}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req request.UpdateProductRequest) (*response.GetProductResponse, *response.ErrorResponse) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	payload := model.Product{
		ID:       id,
		Sku:      req.Sku,
		Name:     req.Name,
		Price:    req.Price,
		IsActive: req.IsActive,
	}
	product, err := s.ProductRepository.Update(ctx, payload)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	result := ProductMapping(product)

	return &response.GetProductResponse{
		Product: result,
	}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) *response.ErrorResponse {
	idProduct, err := uuid.Parse(id)
	if err != nil {
		return response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	err = s.ProductRepository.Delete(ctx, idProduct)
	if err != nil {
		return response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}
	return nil
}

func ProductMapping(product model.Product) response.Product {
	return response.Product{
		ID:       product.ID.String(),
		Sku:      product.Sku,
		Name:     product.Name,
		Price:    product.Price,
		IsActive: product.IsActive,
	}
}
