//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"context"
	"fmt"

	gql_go_practice "github.com/jkieltyka/gql-go-practice"
	"github.com/jkieltyka/gql-go-practice/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

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
		Description:   *transaction.Description,
		Amount:        transaction.Amount,
		ToAccountID:   parties.ToAccount,
		FromAccountID: parties.FromAccount,
	}

	payment, err = createNewPayment(payment)

	return payment, err
}

func (r *mutationResolver) CreateAccount(ctx context.Context, accountInput *models.AccountInput) (*models.Account, error) {
	account := &models.Account{
		Name:       accountInput.Name,
		Balance:    0,
		CustomerID: accountInput.CustomerID,
	}
	account, err := createNewAccount(account)
	return account, err
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, customerInput models.CustomerInput) (*models.Customer, error) {
	customer := &models.Customer{
		FirstName:  customerInput.FirstName,
		LastName:   customerInput.LastName,
		Dob:        customerInput.Dob,
		AccountIDs: make([]string, 0),
	}
	customer, err := createNewCustomer(customer)
	return customer, err
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

	customers, err := fetchCustomers()
	if err != nil {
		return nil, err
	}

	customer, err := findCustomer(customerID, customers)
	if err != nil {
		return nil, err
	}

	accounts, err := fetchAccounts()
	if err != nil {
		return nil, err
	}

	res := make([]*models.Account, 0)

	for _, account := range accounts {
		if account.CustomerID == customer.ID {
			res = append(res, account)
		}
	}

	return res, nil
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
