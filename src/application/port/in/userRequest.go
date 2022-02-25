package in

import (
	"users/src/domain"
)

type UserReqBody struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	IsAmbassador bool   `json:"is_ambassador"`
}

func (userReq *UserReqBody) ToUserDomain() *domain.User {
	return &domain.User{
		FirstName:    userReq.FirstName,
		LastName:     userReq.LastName,
		Email:        userReq.Email,
		Password:     userReq.Password,
		IsAmbassador: userReq.IsAmbassador,
	}
}
