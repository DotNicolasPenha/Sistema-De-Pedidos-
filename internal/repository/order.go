package repository

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order domain.Order) (string, error)
	FindMany() ([]domain.Order, error)
	FindById(id string) (domain.Order, error)
}
