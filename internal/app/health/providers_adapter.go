package health

import (
	"context"
	"errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrMappingProviders  = errors.New("error mapping providers")
	ErrNotFoundProviders = errors.New("not found providers")
)

type adapter struct {
	mysql.Repository
}

func (a *adapter) GetProviders(ctx context.Context, companyID int64) ([]Provider, error) {
	result, err := a.Fetch(ctx, mysql.Statements.Selects.HealthProviders, companyID)
	if err != nil {
		return nil, err
	}

	var providers []Provider
	for result.Next() {
		var provider Provider

		if err := result.Scan(&provider.ID, &provider.CompanyID, &provider.ProviderID, &provider.Name, &provider.CreatedAt, &provider.UpdatedAt); err != nil {
			return nil, ErrMappingProviders
		}

		providers = append(providers, provider)
	}

	if len(providers) == 0 {
		return nil, ErrNotFoundProviders
	}

	return providers, nil
}

func (a *adapter) CreateProvider(ctx context.Context, p Provider) error {
	_, err := a.Save(ctx, mysql.Statements.Inserts.HealthProvider, p.ToString())
	return err
}
