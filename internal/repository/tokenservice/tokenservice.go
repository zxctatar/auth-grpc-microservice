package tokenservice

type TokenService interface {
	Generate(id uint32) string
}
