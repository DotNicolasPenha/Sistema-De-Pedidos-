package service

import (
	"context"
	"errors"

	"com.DotNicolasPenha.SistemaDePedidos/internal/domain"
	"com.DotNicolasPenha.SistemaDePedidos/internal/repository"
)

type PaymentService struct {
	repository repository.PaymentRepository
	ctx        context.Context
}

func (s *PaymentService) Add(payment domain.CreatePaymentDto) (string, error) {
	if payment.Method == "" {
		return "", errors.New("payment method is null")
	}
	if payment.Value == 0 {
		return "", errors.New("payment value is null")
	}
	if _, ok := domain.MethodsPayments[payment.Method]; !ok {
		return "", errors.New("invalid payment method")
	}
	paymentBody := domain.Payment{
		Method: payment.Method,
		Value:  payment.Value,
	}
	return s.repository.Create(s.ctx, paymentBody)
}

func (s *PaymentService) ListAll() ([]domain.Payment, error) {
	return s.repository.FindMany(s.ctx)
}

func (s *PaymentService) ListOneById(id string) (domain.Payment, error) {
	if id == "" {
		return domain.Payment{}, errors.New("id is null")
	}
	return s.repository.FindById(s.ctx, id)
}
