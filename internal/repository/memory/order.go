package memory

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"github.com/google/uuid"
)

type OrderMemoryRepository struct {
	memory []domain.Order
}

func (mr *OrderMemoryRepository) Create(ctx context.Context, order domain.Order) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	order.Id = id
	mr.memory = append(mr.memory, order)
	return order.Id.String(), nil
}

func (mr *OrderMemoryRepository) FindMany() ([]domain.Order, error) {
	return mr.memory, nil
}

func (mr *OrderMemoryRepository) FindById(ctx context.Context, id string) (*domain.Order, error) {
	for _, v := range mr.memory {
		if v.Id.String() == id {
			return &v, nil
		}
	}
	return nil, errors.New("Order not found")
}
