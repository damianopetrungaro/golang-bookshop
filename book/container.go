package book

import (
	"database/sql"
	"github.com/damianopetrungaro/golang-bookshop/book/mapper"
	"github.com/damianopetrungaro/golang-bookshop/book/repository"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"os"
)

func NewContainer() *dig.Container {

	c := dig.New()

	c.Provide(func() (Config, error) {
		sqlUrl := os.Getenv("SQL_CONNECTION_URL")

		return Config{sqlUrl}, nil
	})

	c.Provide(func() (*log.Logger, error) {
		logger := log.New()
		logger.Out = os.Stdout

		return logger, nil
	})

	c.Provide(func(config Config, logger *log.Logger) (*sql.DB, error) {
		db, err := sql.Open("postgres", config.sqlUrl)

		if err != nil {
			logger.Println(err)
			panic(err)
		}

		return db, nil
	})

	c.Provide(func() *mapper.BookMapper {
		return &mapper.BookMapper{}
	})

	c.Provide(func(connection *sql.DB, logger *log.Logger, mapper *mapper.BookMapper) repository.BookRepository {
		return repository.NewBookRepositoryPostgresql(connection, mapper, logger)
	})

	c.Provide(func(connection *sql.DB, logger *log.Logger) repository.AuthorRepository {
		return repository.NewAuthorRepositoryPostgresql(connection, logger)
	})

	return c
}
