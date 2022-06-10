package container

import (
	"users/src/adapter/out/persistence"
	"users/src/utils/config"

	"github.com/jmoiron/sqlx"
)

func Resolve(config config.Config) (Container, error) {
	adapters, err := resolveAdapters(config)
	if err != nil {
		return Container{}, err
	}

	repos, err := resolveRepositories(adapters.Db)

	if err != nil {
		return Container{}, err
	}

	cont := Container{
		Adapters:     adapters,
		Repositories: repos,
	}

	return cont, nil
}

func resolveAdapters(config config.Config) (Adapters, error) {
	psql := persistence.NewDb(config.Database)

	return Adapters{Db: psql}, nil
}

func resolveRepositories(db *sqlx.DB) (Repositories, error) {
	userRepo := persistence.NewPostgresRepository(db)
	repos := Repositories{
		User: userRepo,
	}
	return repos, nil
}
