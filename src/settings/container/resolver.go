package container

import (
	"users/src/adapter/out/persistence"
	"users/src/settings/config"
)

func Resolve(config config.Config) (Container, error) {
	adapters, err := resolveAdapters(config)
	if err != nil {
		return Container{}, err
	}

	repos, err := resolveRepositories("")

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
	return Adapters{}, nil
}

func resolveRepositories(dns string) (Repositories, error) {
	userRepo := persistence.NewPostgresRepository(dns)
	repos := Repositories{
		User: userRepo,
	}
	return repos, nil
}
