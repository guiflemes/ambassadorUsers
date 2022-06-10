package container

import (
	"users/src/application/port/out"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	Adapters     Adapters
	Repositories Repositories
}

type Adapters struct {
	Db *sqlx.DB
}

type Repositories struct {
	User  out.UserRepository
	Login out.LoginRepository
}
