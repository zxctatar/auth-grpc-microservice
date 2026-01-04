package posmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewPostgresUserAuthDataModel_Su(t *testing.T) {
	id := uint32(1)
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("somePass"), bcrypt.DefaultCost)

	posModel := NewPostgresUserAuthDataModel(id, string(hashPassword))

	assert.Equal(t, id, posModel.Id)
	assert.Equal(t, string(hashPassword), posModel.HashPassword)
}
