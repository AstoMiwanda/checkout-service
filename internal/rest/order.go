package rest

import (
	"checkout-service/internal/rest/request"
	"checkout-service/internal/rest/response"
	"checkout-service/pkg/utils"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderService interface {
	CreateOrder(context.Context, request.CreateOrderRequest) (*response.GetOrderResponse, *response.ErrorResponse)
}

type orderHandler struct {
	OrderService OrderService
}

func NewOrderHandler(e *echo.Group, orderService OrderService) {
	handler := &orderHandler{
		OrderService: orderService,
	}

	e.POST("/orders", handler.CreateOrder)

}

func (h *orderHandler) CreateOrder(c echo.Context) error {
	var order request.CreateOrderRequest
	if err := c.Bind(&order); err != nil {
		return utils.JsonErr(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	ctx := c.Request().Context()
	res, err := h.OrderService.CreateOrder(ctx, order)
	if err != nil {
		return utils.JsonErr(c, err.Code, err)
	}

	return utils.JsonOK(c, res)
}
