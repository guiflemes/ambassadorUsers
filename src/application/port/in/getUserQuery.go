package in

type GetUserQuery interface {
	GetAll() ([]*UserRespBody, error)
	GetById(id string) (*UserRespBody, error)
	GetByEmail(email string) (bool, *UserRespBody, error)
}
