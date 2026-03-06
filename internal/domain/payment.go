package domain

import "github.com/google/uuid"

type Payment struct {
	Id     uuid.UUID `json:"id"`
	Method string    `json:"method"`
	Value  int       `json:"value"`
}

type CreatePaymentDto struct {
	Method string `json:"method"`
	Value  int    `json:"value"`
}
