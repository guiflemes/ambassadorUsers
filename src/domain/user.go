package domain

import (
	"time"
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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersList []*User

type Money string

func (u *User) IsValid() (bool, error) {

	err := Validator(u)

	if err != nil {
		return false, err
	}
	return true, nil

}
