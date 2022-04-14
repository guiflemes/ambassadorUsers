package persistence

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const databaseTestName = ""

type postgresTestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *postgresTestSuite) SetupSuite() {
	var err error

	s.db, err = sqlx.Open("postgres", "")
	require.NoError(s.T(), err)

}

func (s *postgresTestSuite) TearDownTest() {
	query := `SELECT TABLE_NAME FROM users.tables WHERE table_schema=´` + databaseTestName + "'"

	rows, err := s.db.Query(query)
	require.NoError(s.T(), err)

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}

		if tableName == "schema_migrations" {
			continue
		}

		queryTruncate := "TRUNCATE TABLE " + tableName
		_, err = s.db.Exec(queryTruncate)
		require.NoError(s.T(), err)
	}
}
