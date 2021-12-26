package company

import (
	"context"
	"errors"
	"fmt"
	oops "github.com/demacedoleo/health-api/internal/app/errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrInternalScanCustomers = errors.New("cannot iterate over rows")
	ErrNotFoundCustomers     = errors.New("not found customers")
)

type customersAdapter struct {
	mysql.Repository
}

func (c customersAdapter) FindCustomers(ctx context.Context, companyID int64) ([]Customer, error) {
	result, err := c.Fetch(ctx, mysql.Statements.Selects.Customers, TypeCustomer.ToString(), companyID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var customers []Customer
	for result.Next() {
		var customer Customer
		if err := result.
			Scan(&customer.Person.ID, &customer.Person.Document, &customer.Person.CompanyID, &customer.Person.Name,
				&customer.Person.LastName, &customer.Person.Birthday, &customer.Person.Gender, &customer.Person.Type,
				&customer.Person.CreatedAt, &customer.Person.UpdatedAt); err != nil {
			return nil, oops.NewError(ErrInternalScan).AddStack(err)
		}

		customers = append(customers, customer)
	}

	if len(customers) == 0 {
		return nil, oops.NewError(ErrNotFound)
	}

	return customers, nil
}

func (c customersAdapter) CreateCustomers(ctx context.Context, ch CustomerHealth) error {
	ch.Person.Type = TypeCustomer.ToString()

	fmt.Println(ch.ToString())

	_, err := c.Save(ctx, mysql.Statements.Inserts.Customer, ch.ToString())
	return err
}

func NewCustomerAdapter(repository mysql.Repository) *customersAdapter {
	return &customersAdapter{repository}
}
