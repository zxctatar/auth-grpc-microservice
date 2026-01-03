package usecaseinterf

import (
	logmodel "auth/internal/usecase/models/login"
	"context"
)

type LoginUseCase interface {
	Login(ctx context.Context, li *logmodel.LoginInput) (string, error)
}
