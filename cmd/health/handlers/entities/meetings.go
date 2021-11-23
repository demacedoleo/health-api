package entities

import "time"

type Meeting struct {
	ID                int64     `json:"id,omitempty"`
	EventID           int64     `json:"event_id"`
	CompanyID         int64     `json:"company_id"`
	Subject           string    `json:"title"`
	MeetStatus        string    `json:"meet_status"`
	StartTime         time.Time `json:"start"`
	EndTime           time.Time `json:"end"`
	AttendantDocument string    `json:"attendant_document,omitempty"`
	AttendantPhone    string    `json:"attendant_phone,omitempty"`
	ResourceID        int64     `json:"resource_id"`
	Modality          string    `json:"modality"`
	CreatedAt         string    `json:"-"`
	UpdatedAt         string    `json:"-"`
}
