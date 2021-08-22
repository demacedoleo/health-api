package test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/demacedoleo/health-api/internal/app/company"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSuiteRoles(t *testing.T) {
	t.Run("Case Init Company ModalitiesAdapter", InitRolesRepository)

	t.Run("Case Success Retrieving Roles", GetRolesSuccess)
	t.Run("Case Empty Result Set Not Found", GetRolesNotFound)
	t.Run("Case Error Fetching Roles", GetRolesRepositoryErr)
	t.Run("Case Error Scan Mapping Roles", GetRolesScanMappingErr)

	t.Run("Case Success Creating Roles", CreateRoleSuccess)
	t.Run("Case Error Creating Roles", CreateRoleError)
}

func InitRolesRepository(t *testing.T) {
	db, _ := NewMock()
	assert.NotNil(t, company.NewRolesAdapter(mysql.NewRepository(db)))
}

func GetRolesSuccess(t *testing.T) {
	expectedRole := company.Role{
		RoleID:    1,
		CompanyID: 2,
		Name:      "Test",
		Type:      "Fake",
	}

	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"id", "company_id", "role_name", "role_type", "create_at", "updated_at"}).
		AddRow(
			expectedRole.RoleID,
			expectedRole.CompanyID,
			expectedRole.Name,
			expectedRole.Type,
			expectedRole.CreatedAt,
			expectedRole.UpdatedAt)

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Roles)).
		ExpectQuery().WithArgs(expectedRole.CompanyID).WillReturnRows(rows)

	r := company.NewRolesAdapter(mysql.NewRepository(db))
	role, err := r.GetRoles(context.Background(), expectedRole.CompanyID)
	assert.Nil(t, err)
	assert.NotNil(t, expectedRole, role)
}

func GetRolesNotFound(t *testing.T) {
	expectedRole := company.Role{
		RoleID:    1,
		CompanyID: 2,
		Name:      "Test",
		Type:      "Fake",
	}

	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"id", "company_id", "role_name", "role_type", "create_at", "updated_at"})

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Roles)).
		ExpectQuery().WithArgs(expectedRole.CompanyID).WillReturnRows(rows)

	r := company.NewRolesAdapter(mysql.NewRepository(db))
	role, err := r.GetRoles(context.Background(), expectedRole.CompanyID)

	assert.Nil(t, role)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, company.ErrNotFoundRoles))
}

func GetRolesRepositoryErr(t *testing.T) {
	expectedID := int64(1)
	expectedRole := make(map[string]interface{})
	expectedRole["fake"] = "fake"
	expectedErr := errors.New("sql err")

	db, mock := NewMock()
	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Roles)).
		ExpectQuery().WithArgs(expectedID).WillReturnError(expectedErr)

	r := company.NewRolesAdapter(mysql.NewRepository(db))
	role, err := r.GetRoles(context.Background(), expectedID)
	assert.Nil(t, role)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, expectedErr))
}

func GetRolesScanMappingErr(t *testing.T) {
	fakeID := "a"
	expectedRole := company.Role{
		RoleID:    1,
		CompanyID: 2,
		Name:      "Test",
		Type:      "Fake",
	}

	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"id", "company_id", "role_name", "role_type", "create_at", "updated_at"}).
		AddRow(
			fakeID,
			expectedRole.CompanyID,
			expectedRole.Name,
			expectedRole.Type,
			expectedRole.CreatedAt,
			expectedRole.UpdatedAt)

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Roles)).
		ExpectQuery().WithArgs(expectedRole.CompanyID).WillReturnRows(rows)

	r := company.NewRolesAdapter(mysql.NewRepository(db))
	role, err := r.GetRoles(context.Background(), expectedRole.CompanyID)

	assert.Nil(t, role)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, company.ErrRoleScan))
}

func CreateRoleSuccess(t *testing.T) {
	expectedRole := company.Role{
		RoleID:    1,
		CompanyID: 2,
		Name:      "Test",
		Type:      "Fake",
	}

	db, mock := NewMock()

	r := company.NewRolesAdapter(mysql.NewRepository(db))

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Inserts.Role)).ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	err := r.CreateRole(context.Background(), expectedRole)
	assert.Nil(t, err)
}

func CreateRoleError(t *testing.T) {
	expectedRole := company.Role{
		RoleID:    1,
		CompanyID: 2,
		Name:      "Test",
		Type:      "Fake",
	}

	expectedErr := errors.New("sql err")

	db, mock := NewMock()
	r := company.NewRolesAdapter(mysql.NewRepository(db))

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Inserts.Role)).ExpectExec().WillReturnError(expectedErr)
	err := r.CreateRole(context.Background(), expectedRole)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, expectedErr))
}
