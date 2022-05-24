package config

import "users/src/settings"

type Database struct {
	User               string
	Password           string
	Host               string
	Port               string
	Database           string
	IdleConnection     int
	OpenConnection     int
	ConnectionLifeTime int
	ConnectionIdleTime int
	ReadTimeout        int
	WriteTimeout       int
	Timeout            int
}

func (db *Database) Parse() {
	db.User = settings.GETENV("POSTGRES_USER")
	db.Password = settings.GETENV("POSTGRES_PASSWORD")
	db.Host = settings.GETENV("POSTGRES_HOST")
	db.Port = settings.GETENV("POSTGRES_PORT")
	db.Database = settings.GETENV("POSTGRES_DB_NAME")
}
