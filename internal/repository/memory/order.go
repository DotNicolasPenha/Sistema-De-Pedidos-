package memory

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"github.com/google/uuid"
)

type OrderMemoryRepository struct {
	data []domain.Order
}

func (mr *OrderMemoryRepository) Create(ctx context.Context, order domain.Order) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	order.Id = id
	mr.data = append(mr.data, order)
	return order.Id.String(), nil
}

func (mr *OrderMemoryRepository) FindMany(ctx context.Context) ([]domain.Order, error) {
	return mr.data, nil
}

func (mr *OrderMemoryRepository) FindById(ctx context.Context, id string) (domain.Order, error) {
	for _, v := range mr.data {
		if v.Id.String() == id {
			return v, nil
		}
	}
	return domain.Order{}, errors.New("Order not found")
}
