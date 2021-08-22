package entities

type Provider struct {
	ID         int64  `json:"id,omitempty"`
	CompanyID  int64  `json:"company_id,omitempty"`
	ProviderID string `json:"provider_id,omitempty"`
	Name       string `json:"provider,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}
