package domain

import (
	"users/src/domain/validators"
)

var (
	Validator = validators.NewValidator().ValidateStruct
)

type User struct {
	Id        string `validate:"required"`
	FirstName string `validate:"required,gt=2"`
	LastName  string `validate:"required,gt=3"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,gt=6"`
	IsActive  bool
}

type UsersList []*User

func (u *User) IsValid() (bool, error) {

	err := Validator(u)

	if err != nil {
		return false, err
	}
	return true, nil

}
