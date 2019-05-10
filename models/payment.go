package models

type Payment struct {
	ID          string `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	ToAccountID   string `json:"toAccountID"`
	FromAccountID string `json:"fromAccountID"`
}