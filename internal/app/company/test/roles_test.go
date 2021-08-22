package test

import (
	"github.com/demacedoleo/health-api/internal/app/company"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRole_ToString(t *testing.T) {
	role := company.Role{
		CompanyID: 1,
		Name:      "Test",
	}

	expectedString := `{"RoleID":0,"CompanyID":1,"Name":"Test","Type":"","CreatedAt":"","UpdatedAt":""}`
	assert.Equal(t, expectedString, role.ToString())
}

func TestRole_ToString_Nil(t *testing.T) {
	var role *company.Role

	expectedString := `null`
	assert.Equal(t, expectedString, role.ToString())
}
