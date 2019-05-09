package models

type Payment struct {
	ID          string
	Description string  
	Amount      float64
	ToAccountID   string
	FromAccountID string
}