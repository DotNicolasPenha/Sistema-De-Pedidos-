package repository

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment domain.Payment) (string, error)
	FindMany(ctx context.Context) ([]domain.Payment, error)
	FindById(ctx context.Context, id string) (domain.Payment, error)
}
