package services

import (
	"users/src/domain"
)

type LoginService interface {
	Authenticate(email string, password string) (bool, *domain.User, error)
}
