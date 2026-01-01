package handler

import (
	regmodels "auth/internal/usecase/models/registration"
	"context"
)

type registrationUCMock struct {
	regUserFn func(ctx context.Context, ri *regmodels.RegInput) (uint32, error)
}

func (r *registrationUCMock) RegUser(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
	return r.regUserFn(ctx, ri)
}
