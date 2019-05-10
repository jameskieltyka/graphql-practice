package models

type Customer struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Dob        string `json:"dob"`
	AccountIDs []string `json:"accountIDs"`
}
