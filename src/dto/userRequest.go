package dto

type UserStoreBodyRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	IsAmbassador bool   `json:"is_ambassador"`
}

type UserLoginBodyRequest struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}
