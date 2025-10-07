package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo OrderRepo
}

func NewService(r OrderRepo) *Service { return &Service{repo: r} }

func (s *Service) CreateOrder(ctx context.Context, customerID string, items []OrderItem) (*Order, error) {

	o := &Order{
		ID:         uuid.NewString(),
		CustomerID: customerID,
		Status:     StatusComplete,
		Items:      items,
		CreatedAt:  time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, o); err != nil {
		return nil, err
	}
	return o, nil
}
