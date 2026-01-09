package handler

import (
	regmodels "auth/internal/usecase/models/registration"
	"context"
)

type registrationUCMock struct {
	regUserFnCalled bool

	regUserFn func(ctx context.Context, ri *regmodels.RegInput) (uint32, error)
}

func (r *registrationUCMock) RegUser(ctx context.Context, ri *regmodels.RegInput) (uint32, error) {
	r.regUserFnCalled = true
	return r.regUserFn(ctx, ri)
}
