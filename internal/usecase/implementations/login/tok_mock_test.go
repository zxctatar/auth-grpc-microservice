package login

type tokenServiceMock struct {
	genFnCalled      bool
	validTokenCalled bool

	genFn        func(id uint32) (string, error)
	validTokenFn func(token string) (uint32, error)
}

func (t *tokenServiceMock) Generate(id uint32) (string, error) {
	t.genFnCalled = true
	return t.genFn(id)
}

func (t *tokenServiceMock) ValidateToken(token string) (uint32, error) {
	t.validTokenCalled = true
	return t.validTokenFn(token)
}
