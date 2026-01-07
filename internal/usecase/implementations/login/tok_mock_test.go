package login

type tokenServiceMock struct {
	genFnCalled      bool
	validTokenCalled bool

	genFn        func(id uint32) (string, error)
	validTokenFn func(token string) (bool, error)
}

func (t *tokenServiceMock) Generate(id uint32) (string, error) {
	t.genFnCalled = true
	return t.genFn(id)
}

func (t *tokenServiceMock) ValidateToken(token string) (bool, error) {
	t.validTokenCalled = true
	return t.validTokenFn(token)
}
