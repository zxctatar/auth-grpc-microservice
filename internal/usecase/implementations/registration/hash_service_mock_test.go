package registration

type hashServiceMock struct {
	genFn func(password []byte) ([]byte, error)
	comFn func(hashPass, password []byte) error
}

func (h *hashServiceMock) Generate(password []byte) ([]byte, error) {
	return h.genFn(password)
}
func (h *hashServiceMock) ComparePassword(hashPass, password []byte) error {
	return h.comFn(hashPass, password)
}
