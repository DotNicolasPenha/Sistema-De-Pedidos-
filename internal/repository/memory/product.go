package memory

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
)

type ProductMemoryRepository struct {
	memory []domain.Product
}

func (mr *ProductMemoryRepository) Create(ctx context.Context, product domain.Product) (string, error) {
	mr.memory = append(mr.memory, product)
	return product.Id.String(), nil
}

func (mr *ProductMemoryRepository) FindMany() ([]domain.Product, error) {
	return mr.memory, nil
}
