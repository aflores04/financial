package request

type CreateTransaction struct {
	Type string `json:"type"`
	Amount float64 `json:"amount`
}