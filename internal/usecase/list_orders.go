package usecase

import (
	"context"

	"github.com/JMKobayashi/Clean-Architeture-Challenge/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]entity.Order, error) {
	orders, err := uc.OrderRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
