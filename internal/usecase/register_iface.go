package usecase

import "context"

type RegistrationUseCase interface {
	RegUser(ctx context.Context, ri *RegInput) (uint32, error)
}
