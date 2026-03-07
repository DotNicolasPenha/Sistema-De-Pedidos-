package memory

import (
	"context"
	"errors"

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

func (mr *ProductMemoryRepository) FindById(ctx context.Context, id string) (*domain.Product, error) {
	for _, v := range mr.memory {
		if v.Id.String() == id {
			return &v, nil
		}
	}
	return nil, errors.New("Product not found")
}
