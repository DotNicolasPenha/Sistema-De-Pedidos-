package service

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"com.DotNicolasPenha.SistemaDePedidos/internal/repository"
)

type ProductService struct {
	repository repository.ProductRepository
	ctx        context.Context
}

func (s *ProductService) Add(product domain.CreateProductDto) (string, error) {
	if product.Category == "" {
		return "", errors.New("category is null")
	}
	if product.Value == 0 {
		return "", errors.New("category value is null")
	}
	if _, ok := domain.Categorys[product.Category]; !ok {
		return "", errors.New("invalid category")
	}
	productBody := domain.Product{
		Category: product.Category,
		Value:    product.Value,
	}
	return s.repository.Create(s.ctx, productBody)
}

func (s *ProductService) ListAll() ([]domain.Product, error) {
	return s.repository.FindMany(s.ctx)
}

func (s *ProductService) ListOneById(id string) (domain.Product, error) {
	if id == "" {
		return domain.Product{}, errors.New("id is null")
	}
	return s.repository.FindById(s.ctx, id)
}
