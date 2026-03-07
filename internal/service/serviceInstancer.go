package service

import (
	"context"

	"com.DotNicolasPenha.SistemaDePedidos/internal/repository"
)

func NewServices(
	ctx context.Context,
	orderRepo repository.OrderRepository,
	paymentRepo repository.PaymentRepository,
	productRepo repository.ProductRepository,
) (OrderService, PaymentService, ProductService) {
	orderService := OrderService{
		repositoryOrder:   orderRepo,
		repositoryProduct: productRepo,
		repositoryPayment: paymentRepo,
		ctx:               ctx,
	}
	paymentService := PaymentService{
		repository: paymentRepo,
		ctx:        ctx,
	}
	productService := ProductService{
		repository: productRepo,
		ctx:        ctx,
	}
	return orderService, paymentService, productService
}
