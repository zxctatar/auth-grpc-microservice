package registration

import (
	userdomain "auth/internal/domain/user"
	"context"
)

type storageRepoMock struct {
	saveFn        func(ctx context.Context, user *userdomain.UserDomain) (uint32, error)
	findByEmailFn func(ctx context.Context, email string) (*userdomain.UserDomain, error)
}

func (m *storageRepoMock) Save(ctx context.Context, user *userdomain.UserDomain) (uint32, error) {
	return m.saveFn(ctx, user)
}

func (m *storageRepoMock) FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error) {
	return m.findByEmailFn(ctx, email)
}
