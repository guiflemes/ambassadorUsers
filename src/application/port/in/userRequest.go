package in

import (
	"users/src/domain"
)

type UserReqBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func ToUserDomain(userReq *UserReqBody) *domain.User {
	return &domain.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Email:     userReq.Email,
		Password:  userReq.Password,
	}
}

type UserUpdateReq struct {
	Id        string `json:"id"  swaggerignore:"true"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserUpdatePasswordDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserUpdateReqToDomain(dto *UserUpdateReq) *domain.User {
	return &domain.User{
		Id:        dto.Id,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}
