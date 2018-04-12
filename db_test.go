package test

import (
	"testing"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func getMockDB() *sqlx.DB {
	db, err := sql.Open("pgmock", "connect_string")
	if err != nil {
		panic(err)
	}

	sqlxDB := sqlx.NewDb(db, "pgmock")
	return sqlxDB
}

func TestGetMockDB(t *testing.T) {
	db := getMockDB()
	defer db.Close()
	assert.NotNil(t, db)
}
