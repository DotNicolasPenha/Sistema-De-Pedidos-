package domain

import "github.com/google/uuid"

type Product struct {
	Id       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Value    int       `json:"value"`
}

type CreateProductDto struct {
	Category string `json:"category"`
	Value    int    `json:"value"`
}
