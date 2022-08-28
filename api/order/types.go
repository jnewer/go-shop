package order

type CompleteOrderRequest struct {
}

type CancelOrderRequest struct {
	OrderId uint `json:"orderId"`
}
