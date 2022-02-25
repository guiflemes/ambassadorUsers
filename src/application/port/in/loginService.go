package in

type LoginService interface {
	Authenticate(email string, password string) (bool, *UserRespBody, error)
}
