package repository

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment domain.Payment) (string, error)
	FindMany() ([]domain.Payment, error)
}
