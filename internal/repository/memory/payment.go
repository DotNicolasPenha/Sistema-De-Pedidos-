package memory

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"github.com/google/uuid"
)

type PaymentMemoryRepository struct {
	data []domain.Payment
}

func (mr *PaymentMemoryRepository) Create(ctx context.Context, payment domain.Payment) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	payment.Id = id
	mr.data = append(mr.data, payment)
	return payment.Id.String(), nil
}

func (mr *PaymentMemoryRepository) FindMany(ctx context.Context) ([]domain.Payment, error) {
	return mr.data, nil
}

func (mr *PaymentMemoryRepository) FindById(ctx context.Context, id string) (domain.Payment, error) {
	for _, v := range mr.data {
		if v.Id.String() == id {
			return v, nil
		}
	}
	return domain.Payment{}, errors.New("Payment not found")
}
