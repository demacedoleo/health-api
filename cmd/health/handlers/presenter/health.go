package presenter

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/internal/app/health"
)

func Providers(providers []health.Provider) []entities.Provider {
	data := make([]entities.Provider, len(providers))
	for i, provider := range providers {
		data[i] = entities.Provider{
			ID:         provider.ID,
			ProviderID: provider.Name,
			Name:       provider.Name,
			CreatedAt:  provider.CreatedAt,
			UpdatedAt:  provider.UpdatedAt,
		}
	}
	return data
}
