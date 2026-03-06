package memory

import (
	"context"

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
