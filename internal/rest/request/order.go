package request

type (
	CreateOrderRequest struct {
		CustomerId string       `json:"customer_id"`
		OrderItems []OrderItems `json:"order_items"`
	}

	OrderItems struct {
		ProductId string `json:"product_id"`
		Qty       int    `json:"qty"`
	}
)
