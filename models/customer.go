package models

type Customer struct {
	ID         string
	FirstName  string
	LastName   string
	Dob        string
	AccountIDs []string
}
