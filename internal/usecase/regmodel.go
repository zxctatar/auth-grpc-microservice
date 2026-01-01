package usecase

type RegInput struct {
	FirstName  string
	MiddleName string
	LastName   string
	Password   string
	Email      string
}

func NewRegInput(firstName, middleName, lastName, password, email string) (*RegInput, error) {
	err := validateRegInput(firstName, lastName, password, email)

	if err != nil {
		return nil, err
	}

	return &RegInput{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Password:   password,
		Email:      email,
	}, nil
}

func validateRegInput(firstName, lastName, password, email string) error {
	if firstName == "" {
		return ErrEmptyFirstName
	}
	if lastName == "" {
		return ErrEmptyLastName
	}
	if password == "" {
		return ErrEmptyPassword
	}
	if email == "" {
		return ErrEmptyEmail
	}
	return nil
}
