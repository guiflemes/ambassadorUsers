package services

import (
	"users/src/dto"
)

type UserService interface {
	GetAll() (*[]dto.UserRespBody, error)
	GetById(id string) (*dto.UserRespBody, error)
	Store(u *dto.UserReqBody) (*dto.UserRespBody, error)
	Update(u *dto.UserReqBody) (*dto.UserRespBody, error)
	Delete(id string) error
}
