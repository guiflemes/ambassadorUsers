package persistence

import (
	"context"
	"fmt"
	"users/src/domain"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresRepository struct {
	client *sqlx.DB
}

func newPostgresSQL(dsn string) *sqlx.DB {

	db, err := sqlx.Connect("postgres", dsn)

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

func (repo *postgresRepository) Store(ctx context.Context, user *domain.User) (*domain.User, error) {

	query := `INSERT into users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)`

	stmt, err := repo.client.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "error preparing stmt to Store an user")
	}

	rows, err := stmt.QueryContext(ctx, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "error storing an user")
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scaning an user after stored")
		}

	}

	return user, nil

}

func (repo *postgresRepository) GetAll(ctx context.Context) (domain.UsersList, error) {
	query := `
		SELECT id, first_name, last_name, email, password, is_active, created_at,
		updated_at FROM users;
	`

	users := &domain.UsersList{}

	err := repo.client.SelectContext(ctx, users, query)
	if err != nil {
		return nil, errors.Wrap(err, "error running select all query db")
	}

	return *users, nil
}

func (repo *postgresRepository) GetBy(ctx context.Context, filter map[string]interface{}) (*domain.User, error) {

	if len(filter) > 1 {
		return nil, errors.New("only one parameter is accepted")
	}

	var field string
	var value interface{}

	for k, v := range filter {
		field, value = k, v
	}

	user := &domain.User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE %s=$1", field)
	err := repo.client.Get(user, query, value)

	if err != nil {
		return nil, errors.Wrap(err, "error getting an user")
	}

	return user, nil

}
