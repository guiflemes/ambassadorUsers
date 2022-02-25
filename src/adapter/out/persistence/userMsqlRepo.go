package persistence

import (
	"users/src/domain"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySQLRepository struct {
	db *gorm.DB
}

func newMySQLDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&domain.User{})

	return db, nil
}

func NewMySQLRepository(dsn string) (*mySQLRepository, error) {
	mySQLDB, err := newMySQLDB(dsn)

	if err != nil {
		return nil, errors.Wrap(err, "DB error")
	}

	repo := &mySQLRepository{
		db: mySQLDB,
	}

	return repo, nil
}
