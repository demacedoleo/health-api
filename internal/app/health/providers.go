package health

import (
	"encoding/json"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
	"log"
)

type Adapter interface {
	Providers
}

type Provider struct {
	ID         int64
	CompanyID  int64
	ProviderID string
	Name       string
	CreatedAt  string
	UpdatedAt  string
}

func (p *Provider) ToString() string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Println("err trying to parse health provider", err)
		return ""
	}

	return string(b)
}

func NewProvidersAdapter(repository mysql.Repository) *adapter {
	return &adapter{repository}
}
