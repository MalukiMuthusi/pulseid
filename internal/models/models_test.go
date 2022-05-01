package models_test

import (
	"testing"

	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomString(t *testing.T) {
	randomStr, err := models.GenerateRandomString()

	assert.Nil(t, err, "error not expected")

	assert.Equal(t, 12, len(randomStr))
}

func TestNewToken(t *testing.T) {
	token, err := models.NewToken()

	assert.Nil(t, err, "error not expected")
	assert.NotNil(t, token, "expected a token structure")
}
