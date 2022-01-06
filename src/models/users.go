package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ModelBase
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"is_ambassador"`
}

func (user *User) SetPassword(pasword string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pasword), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}
