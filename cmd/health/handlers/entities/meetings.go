package entities

import "time"

type Meeting struct {
	ID                int64     `json:"id,omitempty"`
	EventID           int64     `json:"event_id,omitempty"`
	CompanyID         int64     `json:"company_id,omitempty"`
	Subject           string    `json:"subject,omitempty"`
	MeetStatus        string    `json:"meet_status,omitempty"`
	StartTime         time.Time `json:"start_time,omitempty"`
	EndTime           time.Time `json:"end_time,omitempty"`
	ModalityID        int64     `json:"modality_id,omitempty"`
	AttendantDocument string    `json:"attendant_document,omitempty"`
	AttendantPhone    string    `json:"attendant_phone,omitempty"`
	ResourceID        int64     `json:"resource_id,omitempty"`
	CreatedAt         string    `json:"-"`
	UpdatedAt         string    `json:"-"`
}
