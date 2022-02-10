package logic

import (
	"users/src/domain"
	repo "users/src/repositories"
	srv "users/src/services"
)

type userLogic struct {
	userRepo repo.UserRepository
}

func NewUserLogic(userRepo repo.UserRepository) srv.UserService {
	return &userLogic{
		userRepo: userRepo,
	}
}

func (u *userLogic) GetAll() ([]domain.User, error) {
	return u.userRepo.GetAll()
}

func (u *userLogic) GetById(id string) (*domain.User, error) {
	return u.userRepo.GetById(id)
}

func (u *userLogic) Store(data *domain.User) error {
	return u.userRepo.Store(data)
}

func (u *userLogic) Update(id string, data map[string]interface{}) (*domain.User, error) {
	return u.userRepo.Update(id, data)
}

func (u *userLogic) Delete(id string) error {
	return u.userRepo.Delete(id)
}
