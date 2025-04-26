package response

type (
	Order struct {
		CustomerId   string     `json:"customer_id"`
		OrderId      string     `json:"order_id"`
		Total        float64    `json:"total"`
		Discount     float64    `json:"discount"`
		FreeItem     []FreeItem `json:"free_item,omitempty"`
		TotalPayment float64    `json:"total_payment"`
	}

	FreeItem struct {
		ProductId string `json:"product_id"`
		Name      string `json:"name"`
		Qty       int    `json:"qty"`
	}

	GetOrderResponse struct {
		OrderResponse Order `json:"order_response"`
	}
)
