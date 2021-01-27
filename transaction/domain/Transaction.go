package domain

import (
	"github.com/google/uuid"
)

const (
	CREDIT 	= "credit"
	DEBIT	= "debit"
)

type Transaction struct {
	ID uuid.UUID `json:"id"`
	Amount float64 `json:"amount"`
	Type string	`json:"type"`
	EffectiveDate string `json:"effectiveDate"`
}