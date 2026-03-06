package repository

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) (string, error)
	FindMany() ([]domain.Product, error)
	FindById(id string) (domain.Product, error)
}
