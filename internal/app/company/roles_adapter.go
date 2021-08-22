package company

import (
	"context"
	"errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrRoleScan      = errors.New("cannot iterate over rows")
	ErrNotFoundRoles = errors.New("not found roles")
)

type rolesAdapter struct {
	mysql.Repository
}

func (r *rolesAdapter) GetRoles(ctx context.Context, companyID int64) ([]Role, error) {
	result, err := r.Fetch(ctx, mysql.Statements.Selects.Roles, companyID)
	if err != nil {
		return nil, err
	}

	var roles []Role
	for result.Next() {
		var role Role
		if err := result.
			Scan(
				&role.RoleID,
				&role.CompanyID,
				&role.Name,
				&role.Type,
				&role.CreatedAt,
				&role.UpdatedAt); err != nil {
			return nil, ErrRoleScan
		}

		roles = append(roles, role)
	}

	if len(roles) == 0 {
		return nil, ErrNotFoundRoles
	}

	return roles, nil
}

func (r *rolesAdapter) CreateRole(ctx context.Context, role Role) error {
	_, err := r.Save(ctx, mysql.Statements.Inserts.Role, role.ToString())
	return err
}

func NewRolesAdapter(repository mysql.Repository) *rolesAdapter {
	return &rolesAdapter{repository}
}
