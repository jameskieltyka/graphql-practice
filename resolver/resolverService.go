package resolver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jkieltyka/gql-go-practice/models"
)

func matchIds(accountMap map[string]bool, accounts []*models.Account) []*models.Account {
	result := make([]*models.Account, 0)
	for _, acc := range accounts {
		if _, ok := accountMap[acc.ID]; ok {
			result = append(result, acc)
		}
	}
	return result
}

func findAccount(accountID string, accounts []*models.Account) (*models.Account, error) {
	for _, acc := range accounts {
		if acc.ID == accountID {
			return acc, nil
		}

	}
	return nil, fmt.Errorf("could not find account")
}

func findCustomer(customerID string, customers []*models.Customer) (*models.Customer, error) {
	for _, cus := range customers {
		if cus.ID == customerID {
			return cus, nil
		}

	}
	return nil, fmt.Errorf("could not find customer")
}

func findPayments(accountMap map[string]bool, payments []*models.Payment) ([]*models.Payment, error) {
	accPay := make([]*models.Payment, 0)
	for _, pay := range payments {
		if _, ok := accountMap[pay.FromAccountID]; ok {
			accPay = append(accPay, pay)
		}
	}

	return accPay, nil
}

func createNewCustomer(customer *models.Customer) (*models.Customer, error) {
	body, err := json.Marshal(customer)
	res, err := http.Post("http://localhost:3000/customers", "application/json", bytes.NewBuffer(body))
	if err != nil {

	}
	defer res.Body.Close()

	var newCustomer models.Customer
	body, err = ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &newCustomer)
	return &newCustomer, err
}

func createNewAccount(account *models.Account) (*models.Account, error) {
	body, err := json.Marshal(account)
	res, err := http.Post("http://localhost:3000/accounts", "application/json", bytes.NewBuffer(body))
	if err != nil {

	}
	defer res.Body.Close()

	var newAccount models.Account
	body, err = ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &newAccount)
	return &newAccount, err
}

func createNewPayment(payment *models.Payment) (*models.Payment, error) {
	body, err := json.Marshal(payment)
	res, err := http.Post("http://localhost:3000/payments", "application/json", bytes.NewBuffer(body))
	if err != nil {

	}
	defer res.Body.Close()

	var newPayment models.Payment
	body, err = ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &newPayment)
	return &newPayment, err
}

func fetchAccounts() ([]*models.Account, error) {
	res, err := http.Get("http://localhost:3000/accounts")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var allAccounts []*models.Account
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &allAccounts)

	return allAccounts, err
}

func fetchPayments() ([]*models.Payment, error) {
	res, err := http.Get("http://localhost:3000/payments")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var allPayment []*models.Payment
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &allPayment)

	return allPayment, err
}

func fetchCustomers() ([]*models.Customer, error) {
	res, err := http.Get("http://localhost:3000/customers")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var allCustomer []*models.Customer
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &allCustomer)

	return allCustomer, err
}
