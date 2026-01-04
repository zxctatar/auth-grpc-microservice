package login

type hashServiceMock struct {
	genCalled bool
	comCalled bool

	genFn func(password []byte) ([]byte, error)
	comFn func(hashPass, password []byte) error
}

func (h *hashServiceMock) GenerateHashPassword(password []byte) ([]byte, error) {
	h.genCalled = true
	return h.genFn(password)
}

func (h *hashServiceMock) ComparePassword(hashPass, password []byte) error {
	h.comCalled = true
	return h.comFn(hashPass, password)
}
