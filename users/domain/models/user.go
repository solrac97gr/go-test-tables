package models

import "errors"

type User struct {
	ID       int
	Email    string
	Password string
}

var (
	// Validation errors
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	// Repository errors
	ErrSavingUser = errors.New("error saving user")
)

func (u *User) Validate() error {
	if u.Email == "" {
		return ErrInvalidEmail
	}

	if u.Password == "" {
		return ErrInvalidPassword
	}
	return nil
}
