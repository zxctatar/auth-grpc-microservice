package repository

import (
	userdomain "auth/internal/domain/user"
	"context"
)

type StorageRepo interface {
	Save(ctx context.Context, user *userdomain.UserDomain) (uint32, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error)
}
