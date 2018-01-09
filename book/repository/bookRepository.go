package repository

import (
	"errors"
	"github.com/damianopetrungaro/golang-bookshop/book/aggregate"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

var (
	ImpossibleListBooksError    = errors.New("impossible list books")
	ImpossibleRetrieveBookError = errors.New("impossible retrieve book")
	BookNotFoundError           = errors.New("book not found")
	ImpossibleSaveBookError     = errors.New("impossible save book")
	ImpossibleUpdateBookError   = errors.New("impossible update book")
	ImpossibleDeleteBookError   = errors.New("impossible delete book")
)

type BookRepository interface {
	NextId() (valueObject.BookId)
	List(limit int, offset int) ([]aggregate.Book, error)
	Add(aggregate.Book) error
	Get(id valueObject.BookId) (aggregate.Book, error)
	Update(aggregate.Book) error
	Delete(id valueObject.BookId) error
}
