package memory

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type OrderMemoryRepository struct {
	memory []domain.Order
}

func (mr *OrderMemoryRepository) Create(ctx context.Context, order domain.Order) (string, error) {
	mr.memory = append(mr.memory, order)
	return order.Id.String(), nil
}

func (mr *OrderMemoryRepository) FindMany() ([]domain.Order, error) {
	return mr.memory, nil
}
