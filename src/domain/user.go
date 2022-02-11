package domain

type User struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	IsAmbassador bool   `json:"is_ambassador"`
}

type UsersList []*User
