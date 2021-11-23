package company

import (
	"context"
	"errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrMappingCompany  = errors.New("errors mapping company")
	ErrNotFoundCompany = errors.New("not found company")
)

type companyAdapter struct {
	mysql.Repository
}

func (a *companyAdapter) GetCompany(ctx context.Context, id int64) (*Company, error) {
	result, err := a.Fetch(ctx, mysql.Statements.Selects.Company, id)
	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, ErrNotFoundCompany
	}

	var company Company

	if err := result.
		Scan(&company.CompanyID,
			&company.CompanyName,
			&company.CompanyShortName,
			&company.CompanyColor,
			&company.CompanyRegister,
			&company.CreatedAt,
			&company.UpdatedAt,
		); err != nil {
		return nil, ErrMappingCompany
	}

	return &company, nil
}

func (a *companyAdapter) CreateCompany(ctx context.Context, c Company) error {
	_, err := a.Save(ctx, mysql.Statements.Inserts.Company, c.ToString())
	return err
}

func NewCompanyAdapter(repository mysql.Repository) *companyAdapter {
	return &companyAdapter{repository}
}
