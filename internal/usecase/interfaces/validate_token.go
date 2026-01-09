package usecaseinterf

import "context"

type ValidateTokenUseCase interface {
	ValidateToken(ctx context.Context, token string) (uint32, error)
}
