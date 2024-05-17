package models

type Order struct {
	ID       string  `json:"id"`
	Amount   int64   `json:"amount"`
	Currency string  `json:"currency"`
}
