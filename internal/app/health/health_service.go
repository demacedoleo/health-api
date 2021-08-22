package health

import (
	"context"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

type Service interface {
	Providers
}

type Providers interface {
	GetProviders(ctx context.Context, companyID int64) ([]Provider, error)
	CreateProvider(ctx context.Context, provider Provider) error
}

type service struct {
	adapter
}

func (s *service) GetProviders(ctx context.Context, companyID int64) ([]Provider, error) {
	return s.adapter.GetProviders(ctx, companyID)
}

func (s *service) CreateProvider(ctx context.Context, provider Provider) error {
	return s.adapter.CreateProvider(ctx, provider)
}

func NewHealthService(repository mysql.Repository) *service {
	return &service{adapter{repository}}
}
