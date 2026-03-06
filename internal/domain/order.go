package domain

type Order struct {
	Product Product  `json:"product"`
	Payment Payment  `json:"payment"`
	Labels  []string `json:"labels"`
}
