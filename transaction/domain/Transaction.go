package domain

type Transaction struct {
	ID string `json:"id"`
	Amount int `json:"amount"`
	Type string	`json:"type"`
	EffectiveDate string `json:"effectiveDate"`
}