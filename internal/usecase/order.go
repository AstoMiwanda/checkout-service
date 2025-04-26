package usecase

import (
	"checkout-service/internal/repository"
	"checkout-service/internal/rest/request"
	"checkout-service/internal/rest/response"
	"checkout-service/pkg/constant"
	"checkout-service/pkg/model"
	"checkout-service/pkg/utils"
	"context"
	"github.com/google/uuid"
)

type IOrderService interface {
	CreateOrder(context.Context, request.CreateOrderRequest) (*response.GetOrderResponse, *response.ErrorResponse)
}

type OrderService struct {
	OrderRepository    repository.OrderRepository
	DiscountRepository repository.DiscountRepository
	ProductRepository  repository.ProductRepository
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	discountRepository repository.DiscountRepository,
	productRepository repository.ProductRepository,
) *OrderService {
	return &OrderService{
		OrderRepository:    orderRepository,
		DiscountRepository: discountRepository,
		ProductRepository:  productRepository,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req request.CreateOrderRequest) (*response.GetOrderResponse, *response.ErrorResponse) {
	var freeItems []response.FreeItem
	var totalAmount float64
	var discountValue float64

	customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	for _, item := range req.OrderItems {
		productId, err := uuid.Parse(item.ProductId)
		if err != nil {
			continue
		}

		productDetail, err := s.ProductRepository.GetDetail(ctx, productId)
		if err != nil {
			continue
		}
		totalAmount += productDetail.Price * float64(item.Qty)

		discountRules, err := s.DiscountRepository.GetDiscountRuleByProduct(ctx, productId)
		if err != nil {
			continue
		}

		var discountId uuid.UUID
		var getProductId uuid.UUID
		totalBuy := 0
		totalGet := 0
		totalDiscount := 0
		quantityOperator := ""
		for _, discountRule := range discountRules {
			discountId = discountRule.DiscountID
			switch discountRule.Role {
			case constant.BUY:
				totalBuy += discountRule.Quantity
				quantityOperator = discountRule.QuantityOperator
			case constant.GET:
				totalGet += discountRule.Quantity
				getProductId = discountRule.ProductID
			case constant.DISCOUNT:
				totalDiscount += discountRule.Quantity
				getProductId = discountRule.ProductID
			}
		}

		if (quantityOperator == constant.EQUAL && item.Qty == totalBuy) ||
			(quantityOperator == constant.MORE_THAN_EQUAL && item.Qty >= totalBuy) {

			if utils.IsValidUUID(getProductId) {
				product, err := s.ProductRepository.GetDetail(ctx, getProductId)
				if err != nil {
					continue
				}

				if totalGet > 0 {
					freeItems = append(freeItems, response.FreeItem{
						ProductId: product.ID.String(),
						Name:      product.Name,
						Qty:       totalGet,
					})

				} else if totalDiscount > 0 {
					discountValue += product.Price * float64(totalDiscount)
				}

			} else if utils.IsValidUUID(discountId) {
				discount, err := s.DiscountRepository.GetDiscountById(ctx, discountId)
				if err != nil {
					continue
				}
				discountValue += (productDetail.Price * (discount.DiscountValue / 100)) * float64(item.Qty)
			}
		}
	}

	var orderDiscounts []model.OrderDiscount
	for _, freeItem := range freeItems {
		itemProductId, err := uuid.Parse(freeItem.ProductId)
		if err != nil {
			continue
		}
		orderDiscounts = append(orderDiscounts, model.OrderDiscount{
			ProductId: itemProductId,
			Name:      freeItem.Name,
			Qty:       freeItem.Qty,
		})
	}
	payload := model.Order{
		CustomerID:    customerId,
		Status:        constant.OrderStatusPending,
		TotalAmount:   utils.RoundToTwoDecimal(totalAmount),
		TotalDiscount: utils.RoundToTwoDecimal(discountValue),
		TotalPayment:  utils.RoundToTwoDecimal(totalAmount - discountValue),
		OrderDiscount: orderDiscounts,
	}
	order, err := s.OrderRepository.Create(ctx, payload)
	if err != nil {
		return nil, response.HandleError(err, nil, response.DefaultErrorHandlerOptions)
	}

	result := OrderMapping(order, freeItems)

	return &response.GetOrderResponse{
		OrderResponse: result,
	}, nil
}

func OrderMapping(order model.Order, freeItem []response.FreeItem) response.Order {
	return response.Order{
		CustomerId:   order.CustomerID.String(),
		OrderId:      order.ID.String(),
		Total:        order.TotalAmount,
		Discount:     order.TotalDiscount,
		FreeItem:     freeItem,
		TotalPayment: order.TotalPayment,
	}
}
