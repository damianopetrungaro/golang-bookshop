package repository

import (
	"database/sql"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
	log "github.com/sirupsen/logrus"
)

type AuthorRepositoryPostgresql struct {
	connection *sql.DB
	logger     *log.Logger
}

func NewAuthorRepositoryPostgresql(connection *sql.DB, logger *log.Logger) *AuthorRepositoryPostgresql {
	return &AuthorRepositoryPostgresql{connection, logger}
}

func (repo AuthorRepositoryPostgresql) Exists(id valueObject.AuthorId) (bool, error) {
	rows, err := repo.connection.Query("SELECT EXISTS(SELECT id FROM authors WHERE id = $1)", id.Value().String())

	if err != nil {
		repo.logger.Println(err)
		return false, ImpossibleSearchAuthorError
	}

	var exists bool
	rows.Next()
	rows.Scan(&exists)

	return exists, nil
}
