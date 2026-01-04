package hashservice

type HashService interface {
	Generate(password []byte) ([]byte, error)
	ComparePassword(hashPass, password []byte) error
}
