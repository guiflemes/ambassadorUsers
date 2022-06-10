package config

import "users/src/settings"

type App struct {
	Port string
	Host string
}

func (app *App) Parse() {
	app.Port = settings.GETENV("SERVICE_PORT")
	app.Host = settings.GETENV("SERVICE_HOST")
}
