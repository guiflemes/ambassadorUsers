package persistence

import (
	"fmt"
	"testing"
	"users/src/domain"

	"context"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var sHost = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	"testdb",
	"5432",
	"postgres",
	"postgres",
	"testdb",
)

type postgresTestSuite struct {
	suite.Suite
	repo *postgresRepository
}

func (s *postgresTestSuite) SetupSuite() {
	s.repo = NewPostgresRepository(sHost)

}

func (s *postgresTestSuite) seedUsers(users domain.UsersList) {
	query := `INSERT into users (:id, :first_name, :last_name, :email, :password)`
	ctx := context.Background()
	result, err := s.repo.client.NamedExecContext(ctx, query, users)

	require.NoError(s.T(), err)
	rows, err := result.RowsAffected()
	require.NoError(s.T(), err)
	require.True(s.T(), rows > 0)

}

func (s *postgresTestSuite) TestRepoGetAll() {

	users := domain.UsersList{
		&domain.User{
			Id:        "id",
			FirstName: "first",
			LastName:  "last",
			Email:     "email@email.com",
			Password:  "pass123",
			IsActive:  true,
		},
		&domain.User{
			Id:        "id",
			FirstName: "first",
			LastName:  "last",
			Email:     "email@email.com",
			Password:  "pass123",
			IsActive:  true,
		},
	}

	s.seedUsers(users)

	results, err := s.repo.GetAll(context.Background())
	require.NoError(s.T(), err)

	wantID1, wantID2 := results[0].Id, results[1].Id
	ids := []string{users[0].Id, users[1].Id}

	require.Contains(s.T(), ids, wantID1)
	require.Contains(s.T(), ids, wantID2)

}

func TestPostegresRepoSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}
	suite.Run(t, new(postgresTestSuite))
}
