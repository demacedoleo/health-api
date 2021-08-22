package company

import (
	"context"
	"errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrModalitiesScan     = errors.New("cannot iterate over rows")
	ErrNotFoundModalities = errors.New("not found modalities")
)

type modalitiesAdapter struct {
	mysql.Repository
}

func (a *modalitiesAdapter) GetModalities(ctx context.Context, stateID int64) ([]Modality, error) {
	result, err := a.Fetch(ctx, mysql.Statements.Selects.Modalities, stateID)
	if err != nil {
		return nil, err
	}

	var modalities []Modality
	for result.Next() {
		var modality Modality
		if err := result.Scan(&modality.ID, &modality.CompanyID, &modality.Modality, &modality.CreatedAt, &modality.UpdatedAt); err != nil {
			return nil, ErrModalitiesScan
		}

		modalities = append(modalities, modality)
	}

	if len(modalities) == 0 {
		return nil, ErrNotFoundModalities
	}

	return modalities, nil
}

func (a *modalitiesAdapter) CreateModality(ctx context.Context, modality Modality) error {
	_, err := a.Save(ctx, mysql.Statements.Inserts.Modality, modality.ToString())
	return err
}

func NewModalitiesAdapter(repository mysql.Repository) *modalitiesAdapter {
	return &modalitiesAdapter{repository}
}
