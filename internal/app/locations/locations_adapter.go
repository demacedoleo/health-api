package location

import (
	"context"
	"errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrCitiesScan     = errors.New("cannot iterate over rows")
	ErrNotFoundCities = errors.New("not found cities")
)

type adapter struct {
	mysql.Repository
}

func (a *adapter) GetCities(ctx context.Context, stateID int64) ([]City, error) {
	result, err := a.Fetch(ctx, mysql.Statements.Selects.Cities, stateID)
	if err != nil {
		return nil, err
	}

	var cities []City
	for result.Next() {
		var city City
		if err := result.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return nil, ErrCitiesScan
		}

		cities = append(cities, city)
	}

	if len(cities) == 0 {
		return nil, ErrNotFoundCities
	}

	return cities, nil
}

func (a *adapter) GetStates(ctx context.Context) ([]State, error) {
	result, err := a.Fetch(ctx, mysql.Statements.Selects.States)
	if err != nil {
		return nil, err
	}

	var states []State
	for result.Next() {
		var state State
		if err := result.
			Scan(
				&state.ID,
				&state.Name); err != nil {
			return nil, ErrCitiesScan
		}

		states = append(states, state)
	}

	if len(states) == 0 {
		return nil, ErrNotFoundCities
	}

	return states, nil
}

func NewLocationsAdapter(repository mysql.Repository) *adapter {
	return &adapter{
		repository,
	}
}
