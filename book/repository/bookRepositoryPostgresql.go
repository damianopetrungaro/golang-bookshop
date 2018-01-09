package repository

import (
	"database/sql"
	"github.com/damianopetrungaro/golang-bookshop/book/aggregate"
	"github.com/damianopetrungaro/golang-bookshop/book/mapper"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type BookRepositoryPostgresql struct {
	connection *sql.DB
	mapper     *mapper.BookMapper
	logger     *log.Logger
}

func NewBookRepositoryPostgresql(connection *sql.DB, mapper *mapper.BookMapper, logger *log.Logger) *BookRepositoryPostgresql {
	return &BookRepositoryPostgresql{connection, mapper, logger}
}

func (repo BookRepositoryPostgresql) NextId() (valueObject.BookId) {
	id, _ := valueObject.NewBookId(uuid.NewV4().String())

	return *id
}

func (repo BookRepositoryPostgresql) List(limit int, offset int) ([]aggregate.Book, error) {
	var bookSlice []aggregate.Book
	rows, err := repo.connection.Query("SELECT id, title, isbn, year, author_id FROM books ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		repo.logger.Println(err)
		return bookSlice, ImpossibleListBooksError
	}

	for rows.Next() {
		bookSlice = append(bookSlice, *repo.mapper.MapToBookFromRows(rows))
	}

	return bookSlice, nil
}

func (repo BookRepositoryPostgresql) Add(book aggregate.Book) error {
	id, title, isbn, year, authorId := repo.mapper.MapToPlain(book)

	_, err := repo.connection.Exec(
		"INSERT INTO books (id, title, isbn, year, author_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NULL)",
		id, title, isbn, year, authorId,
	)

	if err != nil {
		repo.logger.Println(err)
		return ImpossibleSaveBookError
	}

	return nil
}

func (repo BookRepositoryPostgresql) Get(bookId valueObject.BookId) (aggregate.Book, error) {
	rows, err := repo.connection.Query("SELECT id, title, isbn, year, author_id FROM books WHERE id = $1", bookId.String())

	if err != nil {
		repo.logger.Println(err)
		return aggregate.Book{}, ImpossibleRetrieveBookError
	}

	if !rows.Next() {
		return aggregate.Book{}, BookNotFoundError
	}

	return *repo.mapper.MapToBookFromRows(rows), nil
}

func (repo BookRepositoryPostgresql) Update(book aggregate.Book) error {
	_, err := repo.connection.Exec(
		"UPDATE books SET title = $2, isbn = $3, year = $4, author_id = $5, updated_at = NOW() WHERE id = $1",
		book.Id().String(),
		book.Title().Value(),
		book.Isbn().Value(),
		book.Year().Value(),
		book.AuthorId().String(),
	)

	if err != nil {
		repo.logger.Println(err)
		return ImpossibleUpdateBookError
	}

	return nil
}

func (repo BookRepositoryPostgresql) Delete(id valueObject.BookId) error {
	_, err := repo.connection.Exec("DELETE FROM books WHERE id = $1", id)

	if err != nil {
		repo.logger.Println(err)
		return ImpossibleDeleteBookError
	}

	return nil
}
