package tokenservice

type TokenService interface {
	Generate(id uint32) (string, error)
	ValidateToken(token string) (bool, error)
}
