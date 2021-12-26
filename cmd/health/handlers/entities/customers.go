package entities

import "time"

type Customer struct {
	CompanyID int64  `json:"company_id,omitempty"`
	Person    Person `json:"person"`
	Insurance Health `json:"health"`
}

type Health struct {
	ID             int64     `json:"id,omitempty"`
	Document       string    `json:"document,omitempty"`
	CompanyID      int64     `json:"company_id"`
	HealthName     string    `json:"health_name"`
	MemberID       string    `json:"member_number"`
	MemberStatus   string    `json:"member_status"`
	MemberInitDate string    `json:"member_init_date"`
	MemberEndDate  string    `json:"member_end_date"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
