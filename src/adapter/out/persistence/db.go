package persistence

import (
	"fmt"
	"time"
	"users/src/utils/config"

	"github.com/jmoiron/sqlx"
)

func NewDb(conf config.Database) *sqlx.DB {

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Database,
	)

	db, err := sqlx.Connect("postgres", dns)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(conf.OpenConnection)
	db.SetMaxIdleConns(conf.IdleConnection)
	db.SetConnMaxLifetime(time.Duration(conf.ConnectionLifeTime) * time.Millisecond)
	db.SetConnMaxIdleTime(time.Duration(conf.IdleConnection) * time.Millisecond)

	return db
}
