package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndCompare_Success(t *testing.T) {
	password := []byte("somePass")

	hasher := NewHasher()

	hash, err := hasher.GenerateHashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = hasher.ComparePassword(hash, password)
	assert.NoError(t, err)
}

func TestCompare_WrongPassword(t *testing.T) {
	password := []byte("somePass")
	wrongPassword := []byte("wrongPass")

	hasher := NewHasher()

	hash, err := hasher.GenerateHashPassword(password)
	assert.NoError(t, err)

	err = hasher.ComparePassword(hash, wrongPassword)
	assert.Error(t, err)
}

func TestGenerateHashPassword_DifferentHashes(t *testing.T) {
	password := []byte("somePass")

	hasher := NewHasher()

	hash1, err := hasher.GenerateHashPassword(password)
	assert.NoError(t, err)

	hash2, err := hasher.GenerateHashPassword(password)
	assert.NoError(t, err)

	assert.NotEqual(t, hash1, hash2)
}
