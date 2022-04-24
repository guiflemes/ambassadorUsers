package persistence

import (
	"fmt"
	"log"
	"testing"
	"users/src/domain"

	"context"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var sHost = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	"localhost",
	"5437",
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
	query := `INSERT INTO users (id, first_name, last_name, email, password)
	VALUES (:id, :first_name, :last_name, :email, :password)`
	ctx := context.Background()
	result, err := s.repo.client.NamedExecContext(ctx, query, users)

	require.NoError(s.T(), err)
	rows, err := result.RowsAffected()
	require.NoError(s.T(), err)
	require.True(s.T(), rows > 0)

}

func (s *postgresTestSuite) TestRepoStore() {
	uid1, _ := uuid.NewV4()

	type testCase struct {
		description      string
		user             *domain.User
		expectedErrorMsg string
		idChecker        func(ids ...string)
	}

	for _, scenario := range []testCase{
		{
			description: "Ok insert user with id",
			user: &domain.User{
				Id:        uid1.String(),
				FirstName: "first",
				LastName:  "last",
				Email:     "email@email.com",
				Password:  "pass123",
				IsActive:  true,
			},
			expectedErrorMsg: "",
			idChecker:        func(ids ...string) { s.Equal(ids[0], ids[1]) },
		},
		{
			description: "OK insert user without id",
			user: &domain.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "email2@email.com",
				Password:  "pass123",
				IsActive:  true,
			},
			expectedErrorMsg: "",
			idChecker:        func(ids ...string) { s.NotNil(ids[1]) },
		},
		{
			description: "ERROR insert user with alredy exists email",
			user: &domain.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "email2@email.com",
				Password:  "pass123",
				IsActive:  true,
			},
			expectedErrorMsg: `duplicate key value violates unique constraint "users_email_key"`,
			idChecker:        nil,
		},
	} {
		s.Run(scenario.description, func() {
			result, err := s.repo.Store(context.Background(), scenario.user)

			if err != nil {
				log.Printf("error: %s", err)
				s.Error(err)
				//TODO check error msg
				// s.ErrorContainsf(err )
				return
			}

			s.NotNil(result.CreatedAt)
			s.NotNil(result.UpdatedAt)
			scenario.idChecker(scenario.user.Id, result.Id)

		})
	}

}

func (s *postgresTestSuite) TearDownTest() {
	query := "DELETE FROM users;"
	_, _ = s.repo.client.Query(query)
}

func (s *postgresTestSuite) TestRepoGetAll() {

	uid1, _ := uuid.NewV4()
	uid2, _ := uuid.NewV4()

	users := domain.UsersList{
		&domain.User{
			Id:        uid1.String(),
			FirstName: "first",
			LastName:  "last",
			Email:     "email@email.com",
			Password:  "pass123",
			IsActive:  true,
		},
		&domain.User{
			Id:        uid2.String(),
			FirstName: "first",
			LastName:  "last",
			Email:     "emai2l@email.com",
			Password:  "pass123",
			IsActive:  true,
		},
	}

	s.seedUsers(users)

	results, err := s.repo.GetAll(context.Background())

	s.NoError(err)

	wantID1, wantID2 := results[0].Id, results[1].Id
	ids := []string{users[0].Id, users[1].Id}

	s.Contains(ids, wantID1)
	s.Contains(ids, wantID2)

}

func TestPostegresRepoSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}
	suite.Run(t, new(postgresTestSuite))
}
