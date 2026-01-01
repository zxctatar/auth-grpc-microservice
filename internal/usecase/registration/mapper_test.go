package registration

import (
	"auth/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestModelToDomain_Success(t *testing.T) {
	firstName := "Ivan"
	middleName := "Ivanovich"
	lastName := "Ivanov"
	pass := "somePass"
	email := "mail@mail.ru"

	ri, err := usecase.NewRegInput(
		firstName,
		middleName,
		lastName,
		pass,
		email,
	)

	assert.NoError(t, err)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(ri.Password), bcrypt.DefaultCost)

	assert.NoError(t, err)

	user, err := modelToDomain(ri, string(hashPass))

	assert.NoError(t, err)
	assert.Equal(t, ri.FirstName, user.FirstName)
	assert.Equal(t, ri.MiddleName, user.MiddleName)
	assert.Equal(t, ri.LastName, user.LastName)
	assert.Equal(t, string(hashPass), user.HashPassword)
	assert.Equal(t, ri.Email, user.Email)
}
