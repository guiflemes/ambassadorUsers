package domain

type User struct {
	Id           string `validate: "required"`
	FirstName    string `validate: "required,min=2"`
	LastName     string `validate: "required,min=3"`
	Email        string `validate: "required,email"`
	Password     string `validate: "required,min=6"`
	IsAmbassador bool
}

type UsersList []*User
