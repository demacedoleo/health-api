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

func TestSuiteCompany(t *testing.T) {
	t.Run("Case Init Company ModalitiesAdapter", InitCompanyRepository)

	t.Run("Case Success Retrieving Company", GetCompanySuccess)
	t.Run("Case Error Fetching Company", GetCompanyRepositoryErr)
	t.Run("Case Error Scan Mapping Company", GetCompanyScanMappingErr)
	t.Run("Case Empty Result Set Not Found", GetCompanyNotFound)

	t.Run("Case Success Creating Company", CreateCompanySuccess)
	t.Run("Case Error Retrieving Company", CreateCompanyError)
}

func InitCompanyRepository(t *testing.T) {
	db, _ := NewMock()
	assert.NotNil(t, company.NewCompanyAdapter(mysql.NewRepository(db)))
}

func GetCompanySuccess(t *testing.T) {
	expectedCompany := company.Company{
		CompanyID:        1,
		CompanyName:      "Test Company",
		CompanyShortName: "Test",
		CompanyColor:     "Blue",
		CompanyRegister:  "12345",
	}

	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"company_id", "company_name", "company_short_name", "company_color", "company_register", "create_at", "updated_at"}).
		AddRow(
			expectedCompany.CompanyID,
			expectedCompany.CompanyName,
			expectedCompany.CompanyShortName,
			expectedCompany.CompanyColor,
			expectedCompany.CompanyRegister,
			expectedCompany.CreatedAt,
			expectedCompany.UpdatedAt)

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Company)).
		ExpectQuery().WithArgs(expectedCompany.CompanyID).WillReturnRows(rows)

	a := company.NewCompanyAdapter(mysql.NewRepository(db))
	response, err := a.GetCompany(context.Background(), expectedCompany.CompanyID)
	assert.Nil(t, err)
	assert.NotNil(t, expectedCompany, response)
}

func GetCompanyScanMappingErr(t *testing.T) {
	fakeID := "a"
	expectedCompany := company.Company{
		CompanyID:        1,
		CompanyName:      "Test Company",
		CompanyShortName: "Test",
		CompanyColor:     "Blue",
		CompanyRegister:  "12345",
	}

	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"company_id", "company_name", "company_short_name", "company_color", "company_register", "create_at", "updated_at"}).
		AddRow(
			fakeID,
			expectedCompany.CompanyName,
			expectedCompany.CompanyShortName,
			expectedCompany.CompanyColor,
			expectedCompany.CompanyRegister,
			expectedCompany.CreatedAt,
			expectedCompany.UpdatedAt)

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Company)).
		ExpectQuery().WithArgs(expectedCompany.CompanyID).WillReturnRows(rows)

	a := company.NewCompanyAdapter(mysql.NewRepository(db))
	response, err := a.GetCompany(context.Background(), expectedCompany.CompanyID)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, company.ErrMappingCompany))
}

func GetCompanyRepositoryErr(t *testing.T) {
	expectedID := int64(1)
	expectedCompany := make(map[string]interface{})
	expectedCompany["fake"] = "fake"
	expectedErr := errors.New("sql err")

	db, mock := NewMock()
	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Company)).
		ExpectQuery().WithArgs(expectedID).WillReturnError(expectedErr)

	a := company.NewCompanyAdapter(mysql.NewRepository(db))
	response, err := a.GetCompany(context.Background(), expectedID)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, expectedErr))
}

func GetCompanyNotFound(t *testing.T) {
	expectedCompany := company.Company{
		CompanyID:        1,
		CompanyName:      "Test Company",
		CompanyShortName: "Test",
		CompanyColor:     "Blue",
		CompanyRegister:  "12345",
	}

	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"company_id", "company_name", "company_short_name", "company_color", "company_register", "create_at", "updated_at"})

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Selects.Company)).
		ExpectQuery().WithArgs(expectedCompany.CompanyID).WillReturnRows(rows)

	a := company.NewCompanyAdapter(mysql.NewRepository(db))
	response, err := a.GetCompany(context.Background(), expectedCompany.CompanyID)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, company.ErrNotFoundCompany))
}

func CreateCompanySuccess(t *testing.T) {
	expectedCompany := company.Company{
		CompanyID:        1,
		CompanyName:      "Test Company",
		CompanyShortName: "Test",
		CompanyColor:     "Blue",
		CompanyRegister:  "12345",
	}

	db, mock := NewMock()

	a := company.NewCompanyAdapter(mysql.NewRepository(db))

	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Inserts.Company)).ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	err := a.CreateCompany(context.Background(), expectedCompany)
	assert.Nil(t, err)
}

func CreateCompanyError(t *testing.T) {
	expectedCompany := company.Company{
		CompanyID:        1,
		CompanyName:      "Test Company",
		CompanyShortName: "Test",
		CompanyColor:     "Blue",
		CompanyRegister:  "12345",
	}

	expectedErr := errors.New("sql err")

	db, mock := NewMock()
	mock.ExpectPrepare(regexp.QuoteMeta(mysql.Statements.Inserts.Company)).ExpectExec().WillReturnError(expectedErr)

	a := company.NewCompanyAdapter(mysql.NewRepository(db))

	err := a.CreateCompany(context.Background(), expectedCompany)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, expectedErr))
}
