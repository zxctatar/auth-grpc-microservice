package userdomain

import "net/mail"

type UserDomain struct {
	FirstName    string
	MiddleName   string
	LastName     string
	HashPassword string
	Email        string
}

func NewUserDomain(firstName, middleName, lastName, password, email string) (*UserDomain, error) {
	if !validatePassword(password) {
		return nil, ErrInvalidPassword
	}
	if !validateEmail(email) {
		return nil, ErrInvalidEmail
	}
	return &UserDomain{
		FirstName:    firstName,
		MiddleName:   middleName,
		LastName:     lastName,
		HashPassword: password,
		Email:        email,
	}, nil
}

func validatePassword(password string) bool {
	passwordB := []byte(password)

	if len(passwordB) < 5 || len(passwordB) > 20 {
		return false
	}

	return true
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}
