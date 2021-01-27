package request

type CreateTransaction struct {
	Type string `json:"type" binding:"required,validType"`
	Amount float64 `json:"amount" binding:"required,amount"`
}
