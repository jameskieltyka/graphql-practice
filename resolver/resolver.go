//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	gql_go_practice "github.com/jkieltyka/gql-go-practice"
	"github.com/jkieltyka/gql-go-practice/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var (
	//group of vars to change IDs to a known number
	//only in use for experimentation purposes
	accountIdNum  int
	paymentIdNum  int
	customerIdNum int
)

type Resolver struct {
}

func (r *Resolver) Customer() gql_go_practice.CustomerResolver {
	return &customerResolver{r}
}
func (r *Resolver) Mutation() gql_go_practice.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Payment() gql_go_practice.PaymentResolver {
	return &paymentResolver{r}
}
func (r *Resolver) Query() gql_go_practice.QueryResolver {
	return &queryResolver{r}
}

type customerResolver struct{ *Resolver }

func (r *customerResolver) Accounts(ctx context.Context, obj *models.Customer) ([]*models.Account, error) {
	accountMap := make(map[string]bool)
	for _, id := range obj.AccountIDs {
		accountMap[id] = true
	}

	accounts, err := fetchAccounts()
	if err != nil {
		return nil, err
	}

	return matchIds(accountMap, accounts), nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePayment(ctx context.Context, transaction models.Transaction, parties models.Parties) (*models.Payment, error) {
	paymentIdNum++

	accounts, err := fetchAccounts()
	if err != nil {
		return nil, err
	}

	acc, err := findAccount(parties.ToAccount, accounts)
	if err != nil {
		return nil, fmt.Errorf("could not find the target account")
	}
	acc.Balance += transaction.Amount
	payment := &models.Payment{
		ID:            strconv.Itoa(paymentIdNum),
		Description:   *transaction.Description,
		Amount:        transaction.Amount,
		ToAccountID:   parties.ToAccount,
		FromAccountID: parties.FromAccount,
	}
	// r.payments = append(r.payments, payment)
	return payment, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, accountInput *models.AccountInput) (*models.Account, error) {
	accountIdNum++
	account := &models.Account{
		ID:         strconv.Itoa(accountIdNum),
		Name:       accountInput.Name,
		Balance:    0,
		CustomerID: accountInput.CustomerID,
	}

	// r.accounts = append(r.accounts, account)
	return account, nil
}
func (r *mutationResolver) CreateCustomer(ctx context.Context, customerInput models.CustomerInput) (*models.Customer, error) {
	customerIdNum++
	customer := &models.Customer{
		ID:         strconv.Itoa(customerIdNum),
		FirstName:  customerInput.FirstName,
		LastName:   customerInput.LastName,
		Dob:        customerInput.Dob,
		AccountIDs: make([]string, 0),
	}
	// r.customers = append(r.customers, customer)
	return customer, nil
}

type paymentResolver struct{ *Resolver }

func (r *paymentResolver) ToAccount(ctx context.Context, obj *models.Payment) (*models.Account, error) {
	//TODO replace with better logic later
	accounts, err := fetchAccounts()
	if err != nil {
		return nil, err
	}

	return findAccount(obj.ToAccountID, accounts)
}

func (r *paymentResolver) FromAccount(ctx context.Context, obj *models.Payment) (*models.Account, error) {
	//TODO replace with better logic later
	accounts, err := fetchAccounts()
	if err != nil {
		return nil, err
	}

	return findAccount(obj.FromAccountID, accounts)

}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAccount(ctx context.Context, accountID string) (*models.Account, error) {
	accounts, err := fetchAccounts()
	if err != nil {
		return nil, err
	}

	return findAccount(accountID, accounts)
}
func (r *queryResolver) GetCustomerAccounts(ctx context.Context, customerID string) ([]*models.Account, error) {
	// customer, err := findCustomer(customerID, r.customers)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
func (r *queryResolver) GetCustomerDetails(ctx context.Context, customerID string) (*models.Customer, error) {
	customers, err := fetchCustomers()
	if err != nil {
		return nil, err
	}

	return findCustomer(customerID, customers)
}
func (r *queryResolver) GetCustomerPayments(ctx context.Context, customerID string) ([]*models.Payment, error) {
	customers, err := fetchCustomers()
	if err != nil {
		return nil, err
	}

	custAcc, err := findCustomer(customerID, customers)
	if err != nil {
		return nil, err
	}

	accountMap := make(map[string]bool)
	for _, acc := range custAcc.AccountIDs {
		accountMap[acc] = true
	}

	payments, err := fetchPayments()
	if err != nil {
		return nil, err
	}

	return findPayments(accountMap, payments)
}

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

func fetchAccounts() ([]*models.Account, error) {
	res, err := http.Get("http://localhost:3000/accounts")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var allAccounts *[]*models.Account
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, allAccounts)

	return *allAccounts, err
}

func fetchPayments() ([]*models.Payment, error) {
	res, err := http.Get("http://localhost:3000/payments")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var allPayment *[]*models.Payment
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, allPayment)

	return *allPayment, err
}

func fetchCustomers() ([]*models.Customer, error) {
	res, err := http.Get("http://localhost:3000/customers")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var allCustomer *[]*models.Customer
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, allCustomer)

	return *allCustomer, err
}
