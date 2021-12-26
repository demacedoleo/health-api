package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvironmentDefault(t *testing.T) {
	assert.NotNil(t, "dev", Environment())
}

func TestGetEnvironment(t *testing.T) {
	empty := ""
	prod := "prod"
	assert.Nil(t, os.Setenv("ENV", prod))
	assert.NotNil(t, prod, Environment())
	assert.Nil(t, os.Setenv("ENV", empty))
}
