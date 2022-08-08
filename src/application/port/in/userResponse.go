package in

import (
	_ "users/docs"
	"users/src/domain"
)

type UserRespBody struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewUserRespBody(user *domain.User) *UserRespBody {
	return &UserRespBody{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

type UserDeleteResp struct {
	Id string `json:"id"`
}
