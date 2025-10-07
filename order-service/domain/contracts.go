package domain

import "context"

type OrderRepo interface {
	Create(ctx context.Context, o *Order) error
	SetStatus(ctx context.Context, id, status string) error
	FindByID(ctx context.Context, id string) (*Order, error)
}
