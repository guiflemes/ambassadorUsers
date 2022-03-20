package in

import (
	"users/src/domain"
)

type UserReqBody struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	IsActive bool   `json:"is_ambassador"`
}

func (userReq *UserReqBody) ToUserDomain() *domain.User {
	return &domain.User{
		FirstName:    userReq.FirstName,
		LastName:     userReq.LastName,
		Email:        userReq.Email,
		Password:     userReq.Password,
		IsActive: userReq.IsActive,
	}
}

type UserUpdateReq struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserUpdatePasswordDTO struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

func UserUpdateReqToDomain(dto *UserUpdateReq) *domain.User {
	return &domain.User{
		Id:        dto.Id,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}
