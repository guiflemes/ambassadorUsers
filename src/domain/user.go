package domain

import (
	"users/src/domain/validators"
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
	validator := validators.NewValidator()
	error := validator.ValidateStruct(validator)

	if error != nil {
		return false, error
	}
	return true, nil

}
