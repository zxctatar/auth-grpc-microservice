package login

type tokenServiceMock struct {
	genCalled bool

	genFn func(id uint32) (string, error)
}

func (t *tokenServiceMock) Generate(id uint32) (string, error) {
	t.genCalled = true
	return t.genFn(id)
}
