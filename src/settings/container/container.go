package container

import (
	"users/src/application/port/out"
)

type Container struct {
	Adpaters     Adpaters
	Repositories Repositories
}

type Adpaters struct{}

type Repositories struct {
	User  out.UserRepository
	Login out.LoginRepository
}
