package dto

import "time"

type TransactionDto struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth string
	ExpirationYear  string
	CVV             string
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}
