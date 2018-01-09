package usecase

import (
	"github.com/damianopetrungaro/golang-bookshop/book/aggregate"
	"github.com/damianopetrungaro/golang-bookshop/book/repository"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

func GetBookFromBookshop(bookId valueObject.BookId, bookRepo repository.BookRepository) (aggregate.Book, error) {

	book, err := bookRepo.Get(bookId)

	if err != nil {
		return aggregate.Book{}, err
	}

	return book, nil
}
