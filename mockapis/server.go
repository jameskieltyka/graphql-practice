package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/jkieltyka/gql-go-practice/models"
)

type server struct {
	router    *http.ServeMux
	accounts  []*models.Account
	payments  []*models.Payment
	customers []*models.Customer
}

func main() {
	srv := &server{
		router:    http.NewServeMux(),
		accounts:  make([]*models.Account, 0),
		payments:  make([]*models.Payment, 0),
		customers: make([]*models.Customer, 0),
	}

	srv.router.HandleFunc("/accounts", srv.handleAccounts())
	srv.router.HandleFunc("/payments", srv.handlePayments())
	srv.router.HandleFunc("/customers", srv.handleCustomers())
	
	log.Fatal(http.ListenAndServe(":3000", srv.router))
}

func (s *server) handleAccounts() http.HandlerFunc {
	s.initialiseAccounts()
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := json.Marshal(s.accounts)
		if err != nil {
			
		}
		w.Write(payload)
	}
}

func (s *server) handlePayments() http.HandlerFunc {
	s.initialisePayments()
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := json.Marshal(s.payments)
		if err != nil {
			
		}
		w.Write(payload)
	}
}

func (s *server) handleCustomers() http.HandlerFunc {
	s.initialiseCustomers()
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := json.Marshal(s.customers)
		if err != nil {
			
		}
		w.Write(payload)
	}
}


