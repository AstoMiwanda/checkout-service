package request

type (
	CreateProductRequest struct {
		Sku      string  `json:"sku"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		IsActive bool    `json:"isActive"`
	}

	GetProductListRequest struct {
		Limit int32                  `json:"limit"`
		Page  int32                  `json:"page"`
		Field string                 `json:"field"`
		Sort  string                 `json:"sort"`
		Where map[string]interface{} `json:"where"`
	}

	UpdateProductRequest struct {
		ID       string  `json:"Id"`
		Sku      string  `json:"sku"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		IsActive bool    `json:"isActive"`
	}
)
