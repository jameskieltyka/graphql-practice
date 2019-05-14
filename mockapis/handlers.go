package main

import(
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/jkieltyka/gql-go-practice/models"
	"fmt"
	"strconv"
)

func (s *server) handleAccounts() http.HandlerFunc {
	s.initialiseAccounts()
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: 
			payload, err := json.Marshal(s.accounts)
			if err != nil {
			
			}
			w.Write(payload)
		case http.MethodPost:
			var newAccount models.Account
			accounts, err := ioutil.ReadAll(r.Body)
			if err != nil {
			
			}
			err = json.Unmarshal(accounts, &newAccount)
			if err != nil {
			
			}
			newAccount.ID = strconv.Itoa(len(s.accounts) + 1)
			s.accounts = append(s.accounts, &newAccount)
			
			payload, err := json.Marshal(newAccount)
			if err != nil {
				
			}
			w.Write(payload)
		}
	}
}

func (s *server) handlePayments() http.HandlerFunc {
	s.initialisePayments()
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet: 
				payload, err := json.Marshal(s.payments)
				if err != nil {
					
				}
				w.Write(payload)
			case http.MethodPost:
				var newPayment models.Payment
				payment, err := ioutil.ReadAll(r.Body)
				if err != nil {
				
				}
				err = json.Unmarshal(payment, &newPayment)
				newPayment.ID = strconv.Itoa(len(s.payments) + 1)
				if err != nil {
				
				}
				s.payments = append(s.payments, &newPayment)

				payload, err := json.Marshal(newPayment)
				if err != nil {
					
				}
				w.Write(payload)
			}
	}
}

func (s *server) handleCustomers() http.HandlerFunc {
	s.initialiseCustomers()
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet: 
				payload, err := json.Marshal(s.customers)
				if err != nil {
					
				}
				w.Write(payload)
			case http.MethodPost:
				var newCustomer models.Customer
				customer, err := ioutil.ReadAll(r.Body)
				if err != nil {
			
				}
				err = json.Unmarshal(customer, &newCustomer)
				newCustomer.ID = strconv.Itoa(len(s.customers) + 1)
				fmt.Println(newCustomer.ID)
				if err != nil {
				
				}
				s.customers = append(s.customers, &newCustomer)

				payload, err := json.Marshal(newCustomer)
				if err != nil {
					
				}
				w.Write(payload)
			}
	}
}
