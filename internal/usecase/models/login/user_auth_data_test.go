package logmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUserAuthData_Success(t *testing.T) {
	id := uint32(1)
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("somePass"), bcrypt.DefaultCost)

	authData := NewUserAuthData(id, string(hashPassword))

	assert.Equal(t, authData.Id, id)
	assert.Equal(t, authData.HashPassword, string(hashPassword))
}
