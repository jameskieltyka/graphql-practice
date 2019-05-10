package main

import (
	"github.com/jkieltyka/gql-go-practice/models"
)

func (s *server) initialiseCustomers() {
	s.customers = append(s.customers,
		&models.Customer{
			ID: "1",
			FirstName: "test1",
			LastName: "lasttest",
			Dob: "21/11/78",
			AccountIDs: []string{"1","3"},
		},
		&models.Customer{
			ID: "2",
			FirstName: "test 2",
			LastName: "lasttest",
			Dob: "11/02/82",
			AccountIDs:  []string{"2"},
		},
	)
}

func (s *server) initialiseAccounts() {
	s.accounts = append(s.accounts,
		&models.Account{
			ID: "1",
			Name: "test 1",
			Balance: 2131.12,
			CustomerID: "1",
		},
		&models.Account{
			ID: "2",
			Name: "test 2",
			Balance: 153.22,
			CustomerID: "2",
		},
		&models.Account{
			ID: "3",
			Name: "test3",
			Balance: 6767.12,
			CustomerID: "1",
		},
	)
}

func (s *server) initialisePayments() {
	s.payments = append(s.payments,
		&models.Payment{
			ID: "1",     
			Description: "Pay test 2",
			Amount: 12.11,      
			ToAccountID: "2",  
			FromAccountID: "3", 
		},
	)
}
