package domain

import (
	"time"
	"users/src/domain/validators"
)

var (
	Validator = validators.NewValidator().ValidateStruct
)

type User struct {
	Id        string    `validate:"required" db:"id"`
	FirstName string    `validate:"required,gt=2" db:"first_name"`
	LastName  string    `validate:"required,gt=3" db:"last_name"`
	Email     string    `validate:"required,email" db:"email"`
	Password  string    `validate:"required,gt=6" db:"password"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
