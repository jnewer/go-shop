package cart

type ItemCartRequest struct {
	SKU   string `json:"sku"`
	Count int    `json:"count"`
}

type CreateCategoryResponse struct {
	Message string `json:"message"`
}
