package userdomain

import "net/mail"

type UserDomain struct {
	FirstName string
	MiddleName string
	LastName string
	Password string
	Email string
}

func NewUserDomain(firstName, middleName, lastName, password, email string) (*UserDomain, error) {
	if !validateEmail(email) {
		return nil, ErrInvalidEmail
	}
	return &UserDomain{
		FirstName: firstName,
		MiddleName: middleName,
		LastName: lastName,
		Password: password,
		Email: email,
	}, nil
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}