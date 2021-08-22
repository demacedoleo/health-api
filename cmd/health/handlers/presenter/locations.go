package presenter

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	location "github.com/demacedoleo/health-api/internal/app/locations"
)

func States(states []location.State) []entities.State {
	output := make([]entities.State, len(states))
	for i, state := range states {
		output[i] = entities.State{
			ID:   state.ID,
			Name: state.Name,
		}
	}
	return output
}

func Cities(cities []location.City) []entities.City {
	output := make([]entities.City, len(cities))
	for i, city := range cities {
		output[i] = entities.City{
			ID:      city.ID,
			StateID: city.StateID,
			Name:    city.Name,
		}
	}
	return output
}
