package test

import (
	company2 "github.com/demacedoleo/health-api/internal/app/company"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompany_ToString(t *testing.T) {
	company := company2.Company{
		CompanyID:   1,
		CompanyName: "Test",
	}

	expectedString := `{"CompanyID":1,"CompanyName":"Test","CompanyShortName":"","CompanyColor":"","CompanyRegister":"","CreatedAt":"","UpdatedAt":""}`
	assert.Equal(t, expectedString, company.ToString())
}

func TestCompany_ToString_Nil(t *testing.T) {
	var company *company2.Company

	expectedString := `null`
	assert.Equal(t, expectedString, company.ToString())
}
