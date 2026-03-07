package service

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"com.DotNicolasPenha.SistemaDePedidos/internal/repository"
)

type OrderService struct {
	repositoryOrder   repository.OrderRepository
	repositoryProduct repository.ProductRepository
	repositoryPayment repository.PaymentRepository
	ctx               context.Context
}

func (s *OrderService) Add(order domain.CreateOrderDto) (string, error) {
	if order.PaymentID == "" {
		return "", errors.New("Order Payment ID is null")
	}
	if order.ProductID == "" {
		return "", errors.New("Order Product ID is null")
	}
	product, err := s.repositoryProduct.FindById(s.ctx, order.ProductID)
	if err != nil {
		return "", err
	}
	payment, err := s.repositoryPayment.FindById(s.ctx, order.PaymentID)
	if err != nil {
		return "", err
	}
	if product.Value != payment.Value {
		return "", errors.New("Product value and Payment value not match")
	}
	newOrderBody := domain.Order{
		Labels:  renderOrderLabels(product),
		Product: product,
		Payment: payment,
	}
	renderConditionsPayment(newOrderBody.Payment)
	orderId, err := s.repositoryOrder.Create(s.ctx, newOrderBody)
	return orderId, err
}

func (s *OrderService) ListAll() ([]domain.Order, error) {
	return s.repositoryOrder.FindMany(s.ctx)
}

func (s *OrderService) ListOneById(id string) (domain.Order, error) {
	if id == "" {
		return domain.Order{}, errors.New("Id is null")
	}
	return s.repositoryOrder.FindById(s.ctx, id)
}
func renderConditionsPayment(payment domain.Payment) {
	if payment.Method == domain.MethodsPayments["Boleto"] {
		// discount 10%
		discount := 10
		payment.Value -= payment.Value / discount
	}
}

func renderOrderLabels(product domain.Product) []string {
	labels := []string{}
	if label, ok := domain.CategoryLabels[product.Category]; ok {
		labels = append(labels, label)
	}
	if product.Value > 1000 {
		labels = append(labels, domain.Labels["FreteGrátis"])
	}
	return labels
}
