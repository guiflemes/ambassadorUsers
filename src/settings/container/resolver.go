package container

import (
	"users/src/adapter/out/persistence"
	"users/src/settings/config"
)

func Resolve(config config.Config) (Container, error) {

}

func resolveRepositories(dns string) (Repositories, error) {
	userRepo := persistence.NewPostgresRepository(dns)
	repos := Repositories{
		User: userRepo,
	}
	return repos, nil
}
