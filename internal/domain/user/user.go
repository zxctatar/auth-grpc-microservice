package userdomain

import "net/mail"

type UserDomain struct {
	FirstName    string
	MiddleName   string
	LastName     string
	HashPassword string
	Email        string
}

func NewUserDomain(firstName, middleName, lastName, hashPassword, email string) (*UserDomain, error) {
	if !validateEmail(email) {
		return nil, ErrInvalidEmail
	}
	return &UserDomain{
		FirstName:    firstName,
		MiddleName:   middleName,
		LastName:     lastName,
		HashPassword: hashPassword,
		Email:        email,
	}, nil
}

func RestoreUserDomain(firstName, middleName, lastName, hashPassword, email string) *UserDomain {
	return &UserDomain{
		FirstName:    firstName,
		MiddleName:   middleName,
		LastName:     lastName,
		HashPassword: hashPassword,
		Email:        email,
	}
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}
