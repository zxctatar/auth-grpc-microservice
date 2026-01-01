package usecaseinterf

import (
	regmodels "auth/internal/usecase/models/registration"
	"context"
)

type RegistrationUseCase interface {
	RegUser(ctx context.Context, ri *regmodels.RegInput) (uint32, error)
}
