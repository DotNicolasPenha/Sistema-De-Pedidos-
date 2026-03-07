package repository

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order domain.Order) (string, error)
	FindMany(ctx context.Context) ([]domain.Order, error)
	FindById(ctx context.Context, id string) (domain.Order, error)
}
