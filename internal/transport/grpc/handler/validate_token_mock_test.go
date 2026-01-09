package handler

import "context"

type validateTokenUCMock struct {
	validateFnCalled bool

	validateFn func(ctx context.Context, token string) (uint32, error)
}

func (v *validateTokenUCMock) ValidateToken(ctx context.Context, token string) (uint32, error) {
	v.validateFnCalled = true
	return v.validateFn(ctx, token)
}
