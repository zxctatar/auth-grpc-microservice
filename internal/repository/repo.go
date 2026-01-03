package repository

import (
	userdomain "auth/internal/domain/user"
	logmodel "auth/internal/usecase/models/login"
	"context"
)

type StorageRepo interface {
	Save(ctx context.Context, user *userdomain.UserDomain) (uint32, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error)
	FindAuthDataByEmail(ctx context.Context, email string) (*logmodel.UserAuthData, error)
}
