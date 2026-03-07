package memory

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"github.com/google/uuid"
)

type ProductMemoryRepository struct {
	data []domain.Product
}

func (mr *ProductMemoryRepository) Create(ctx context.Context, product domain.Product) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	product.Id = id
	mr.data = append(mr.data, product)
	return product.Id.String(), nil
}

func (mr *ProductMemoryRepository) FindMany(ctx context.Context) ([]domain.Product, error) {
	return mr.data, nil
}

func (mr *ProductMemoryRepository) FindById(ctx context.Context, id string) (domain.Product, error) {
	for _, v := range mr.data {
		if v.Id.String() == id {
			return v, nil
		}
	}
	return domain.Product{}, errors.New("Product not found")
}
