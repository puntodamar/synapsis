package http

import (
	"github.com/synapsis/order-service/domain"
)

type CreateOrderDTO struct {
	CustomerID string         `json:"customer_id"`
	Items      []OrderItemDTO `json:"items"`
}

type CreateOrderResponse struct {
	OrderID string        `json:"order_id"`
	Status  domain.Status `json:"status"`
}

type OrderItemDTO struct {
	SKU string `json:"sku"`
	Qty int32  `json:"qty"`
}
