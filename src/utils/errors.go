package utils

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserInvalid       = errors.New("user invalid")
	ErrUserAlreadyExists = errors.New("user with the given email already exists")
	ErrInvalidParameter  = errors.New("invalid parameter")
	ErrUnauthorized      = errors.New("user unauthorized")
)
