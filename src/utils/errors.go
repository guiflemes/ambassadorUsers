package utils

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserInvalid      = errors.New("user invalid")
	ErrUserAlredyExists = errors.New("user already exists")
)
