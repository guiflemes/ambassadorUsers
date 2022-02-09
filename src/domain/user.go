package domain

type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"is_ambassador"`
}
