package response

type (
	Product struct {
		ID       string  `json:"Id"`
		Sku      string  `json:"sku"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		IsActive bool    `json:"isActive"`
	}

	GetProductResponse struct {
		Product Product `json:"product"`
	}

	GetProductListResponse struct {
		Products  []Product `json:"products"`
		Page      int32     `json:"page"`
		Limit     int32     `json:"limit"`
		Total     int64     `json:"total"`
		TotalPage int64     `json:"totalPage"`
	}
)
