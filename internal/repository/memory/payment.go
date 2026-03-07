package memory

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type PaymentMemoryRepository struct {
	memory []domain.Payment
}

func (mr *PaymentMemoryRepository) Create(ctx context.Context, payment domain.Payment) (string, error) {
	mr.memory = append(mr.memory, payment)
	return payment.Id.String(), nil
}

func (mr *PaymentMemoryRepository) FindMany() ([]domain.Payment, error) {
	return mr.memory, nil
}

func (mr *PaymentMemoryRepository) FindById(ctx context.Context, id string) (*domain.Payment, error) {
	for _, v := range mr.memory {
		if v.Id.String() == id {
			return &v, nil
		}
	}
	return nil, errors.New("Payment not found")
}
