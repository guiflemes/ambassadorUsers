package persistence

import (
	"context"
	"users/src/domain"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
)

type postgresRepository struct {
	client *sqlx.DB
}

func newPostgresSQL(dsn string) *sqlx.DB {

	db, err := sqlx.Connect("postgres", dsn)
	log.Println(err)

	if err != nil {
		panic(err)
	}
	return db
}

func NewPostgresRepository(dsn string) *postgresRepository {

	postgresDB := newPostgresSQL(dsn)

	repo := &postgresRepository{
		client: postgresDB,
	}

	return repo
}

func (repo *postgresRepository) GetAll(ctx context.Context) (domain.UsersList, error) {
	query := `
		SELECT id, first_name, last_name, email, password, is_active, created_at,
		updated_at FROM users;
	`

	stmt, err := repo.client.PrepareContext(ctx, query)

	if err != nil {
		return nil, errors.Wrap(err, "repository error trying to prepare stmt")
	}

	rows, err := stmt.QueryContext(ctx)

	defer rows.Close()

	if err != nil {
		return nil, errors.Wrap(err, "repository error trying to query all users")
	}

	var users domain.UsersList

	for rows.Next() {
		user := &domain.User{}

		err := rows.Scan(
			&user.Id,
			&user.FirstName,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		users = append(users, user)

	}

	return users, nil
}
