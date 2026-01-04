package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Generate(password []byte) ([]byte, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashPass, err
}

func (h *Hasher) ComparePassword(hashPass, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPass, password)
	return err
}
