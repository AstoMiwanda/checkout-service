package rest

import (
	"checkout-service/internal/rest/request"
	"checkout-service/internal/rest/response"
	"checkout-service/pkg/utils"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductService interface {
	CreateProduct(context.Context, request.CreateProductRequest) (*response.GetProductResponse, *response.ErrorResponse)
	GetDetailProduct(context.Context, string) (*response.GetProductResponse, *response.ErrorResponse)
	GetListProduct(context.Context, request.GetProductListRequest) (*response.GetProductListResponse, *response.ErrorResponse)
	UpdateProduct(context.Context, request.UpdateProductRequest) (*response.GetProductResponse, *response.ErrorResponse)
	DeleteProduct(context.Context, string) *response.ErrorResponse
}

type productHandler struct {
	ProductService ProductService
}

func NewProductHandler(e *echo.Group, productService ProductService) {
	handler := &productHandler{
		ProductService: productService,
	}

	e.POST("/products", handler.CreateProduct)
	e.GET("/products/:id", handler.GetDetailProduct)
	e.GET("/products", handler.GetListProduct)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)

}

func (h *productHandler) CreateProduct(c echo.Context) error {
	var product request.CreateProductRequest
	if err := c.Bind(&product); err != nil {
		return utils.JsonErr(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	ctx := c.Request().Context()
	res, err := h.ProductService.CreateProduct(ctx, product)
	if err != nil {
		return utils.JsonErr(c, err.Code, err)
	}

	return utils.JsonOK(c, res)
}

func (h *productHandler) GetDetailProduct(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	res, err := h.ProductService.GetDetailProduct(ctx, id)
	if err != nil {
		return utils.JsonErr(c, err.Code, err)
	}

	return utils.JsonOK(c, res)
}

func (h *productHandler) GetListProduct(c echo.Context) error {
	var req request.GetProductListRequest
	if err := c.Bind(&req); err != nil {
		return utils.JsonErr(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	ctx := c.Request().Context()
	res, err := h.ProductService.GetListProduct(ctx, req)
	if err != nil {
		return utils.JsonErr(c, err.Code, err)
	}

	return utils.JsonOK(c, res)
}

func (h *productHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product request.UpdateProductRequest
	if err := c.Bind(&product); err != nil {
		return utils.JsonErr(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	product.ID = id
	ctx := c.Request().Context()
	res, err := h.ProductService.UpdateProduct(ctx, product)
	if err != nil {
		return utils.JsonErr(c, err.Code, err)
	}

	return utils.JsonOK(c, res)
}

func (h *productHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	err := h.ProductService.DeleteProduct(ctx, id)
	if err != nil {
		return utils.JsonErr(c, err.Code, err)
	}

	return utils.JsonOK(c, response.BaseResponse{
		Code:    fmt.Sprintf("%d", http.StatusOK),
		Message: "Successfully",
	})
}
