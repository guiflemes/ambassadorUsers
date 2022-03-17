package in

type UserService interface {
	GetAll() ([]*UserRespBody, error)
	GetById(id string) (*UserRespBody, error)
	Store(u *UserReqBody) (*UserRespBody, error)
	Update(u *UserUpdateReq) (*UserRespBody, error)
	Delete(id string) error
}
