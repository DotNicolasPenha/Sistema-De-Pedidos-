package memory

import "com.DotNicolasPenha.SistemaDePedidos/internal/domain"

func NewMemoryRepositorys() (*OrderMemoryRepository, *PaymentMemoryRepository, *ProductMemoryRepository) {
	orderMemoryRepository := OrderMemoryRepository{data: []domain.Order{}}
	paymentMemoryRepository := PaymentMemoryRepository{data: []domain.Payment{}}
	productMemoryRepository := ProductMemoryRepository{data: []domain.Product{}}

	return &orderMemoryRepository, &paymentMemoryRepository, &productMemoryRepository
}
