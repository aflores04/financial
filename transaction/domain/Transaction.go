package domain

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID uuid.UUID `json:"id"`
	Amount float64 `json:"amount"`
	Type string	`json:"type"`
	EffectiveDate string `json:"effectiveDate"`
}