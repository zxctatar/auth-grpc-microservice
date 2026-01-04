package hashservice

type HashService interface {
	GenerateHashPassword(password []byte) ([]byte, error)
	ComparePassword(hashPass, password []byte) error
}
