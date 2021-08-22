package location

import (
	"context"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

type LocationsService interface {
	GetCities(ctx context.Context, stateID int64) ([]City, error)
	GetStates(ctx context.Context) ([]State, error)
}

type locationsService struct {
	adapter
}

func (l *locationsService) GetCities(ctx context.Context, stateID int64) ([]City, error) {
	return l.adapter.GetCities(ctx, stateID)
}

func (l *locationsService) GetStates(ctx context.Context) ([]State, error) {
	return l.adapter.GetStates(ctx)
}

func NewLocationsService(repository mysql.Repository) *locationsService {
	return &locationsService{
		adapter{repository},
	}
}
