package logmodel

type LoginInput struct {
	Email    string
	Password string
}

func NewLoginInput(email, password string) (*LoginInput, error) {
	err := validateLoginInput(email, password)

	if err != nil {
		return nil, err
	}

	return &LoginInput{
		Email:    email,
		Password: password,
	}, nil
}

func validateLoginInput(email, password string) error {
	if email == "" {
		return ErrEmptyEmail
	}
	if password == "" {
		return ErrEmptyPassword
	}
	return nil
}

type UserAuthData struct {
	Id           uint32
	HashPassword string
}

func NewUserAuthData(id uint32, hashPassword string) *UserAuthData {
	return &UserAuthData{
		Id:           id,
		HashPassword: hashPassword,
	}
}
