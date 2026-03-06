package domain

type Payment struct {
	Method string `json:"method"`
	Value  int    `json:"value"`
}
