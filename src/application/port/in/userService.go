package in

type UserService interface {
	GetAll() ([]*UserRespBody, error)
	GetById(id string) (*UserRespBody, error)
	Store(u *UserReqBody) (*UserRespBody, error)
	Update(u *UserUpdateDTO) (*UserRespBody, error)
	Delete(id string) error
}
