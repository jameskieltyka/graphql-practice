package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/jkieltyka/gql-go-practice/models"
)

type server struct {
	router    *mux.Router
	accounts  []*models.Account
	payments  []*models.Payment
	customers []*models.Customer
}

func main() {
	srv := &server{
		router:    mux.NewRouter(),
		accounts:  make([]*models.Account, 0),
		payments:  make([]*models.Payment, 0),
		customers: make([]*models.Customer, 0),
	}

	srv.router.Handle("/accounts", srv.handleAccounts()).Methods(http.MethodGet, http.MethodPost)
	srv.router.Handle("/payments", srv.handlePayments()).Methods(http.MethodGet, http.MethodPost)
	srv.router.Handle("/customers", srv.handleCustomers()).Methods(http.MethodGet, http.MethodPost)

	log.Fatal(http.ListenAndServe(":3000", srv.router))
}


