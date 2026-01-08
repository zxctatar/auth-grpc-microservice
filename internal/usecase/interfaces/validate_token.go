package usecaseinterf

type ValidateTokenUseCase interface {
	ValidateToken(token string) (uint32, error)
}
