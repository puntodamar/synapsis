package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/synapsis/order-service/domain"
	"github.com/synapsis/order-service/transport/http/response"
)

type Handler struct {
	service *domain.Service
}

func NewHandler(s *domain.Service) *Handler { return &Handler{service: s} }

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var in CreateOrderDTO
	if err := c.BodyParser(&in); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}
	if in.CustomerID == "" || len(in.Items) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "customer_id and items required")
	}

	items := make([]domain.OrderItem, 0, len(in.Items))
	for _, it := range in.Items {
		if it.SKU == "" || it.Qty <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "each item needs sku and qty>0")
		}
		items = append(items, domain.OrderItem{SKU: it.SKU, Qty: it.Qty})
	}

	o, err := h.service.CreateOrder(c.Context(), in.CustomerID, items)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, response.Envelope{
		"data": CreateOrderResponse{
			OrderID: o.ID,
			Status:  o.Status,
		},
	})
}
