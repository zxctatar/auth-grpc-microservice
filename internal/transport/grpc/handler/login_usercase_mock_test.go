package handler

import (
	logmodel "auth/internal/usecase/models/login"
	"context"
)

type loginUCMock struct {
	logFnCalled bool

	logFn func(ctx context.Context, li *logmodel.LoginInput) (string, error)
}

func (l *loginUCMock) Login(ctx context.Context, li *logmodel.LoginInput) (string, error) {
	l.logFnCalled = true
	return l.logFn(ctx, li)
}
