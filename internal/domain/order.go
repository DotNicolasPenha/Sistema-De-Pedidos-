package domain

import "github.com/google/uuid"

type Order struct {
	Id      uuid.UUID `json:"id"`
	Product Product   `json:"product"`
	Payment Payment   `json:"payment"`
	Labels  []string  `json:"labels"`
}
