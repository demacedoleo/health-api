package company

import (
	"encoding/json"
	"log"
)

type ModalitiesAdapter interface {
	Modalities
}

type Modality struct {
	ID        int64
	CompanyID int64
	Modality  string
	CreatedAt string
	UpdatedAt string
}

func (m *Modality) ToString() string {
	b, err := json.Marshal(m)
	if err != nil {
		log.Println("err trying to parse modality", err)
		return ""
	}

	return string(b)
}
