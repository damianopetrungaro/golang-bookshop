package usecase

import (
	"github.com/damianopetrungaro/golang-bookshop/book/aggregate"
	"github.com/damianopetrungaro/golang-bookshop/book/repository"
	"github.com/damianopetrungaro/golang-bookshop/shared"
)

func ListBooksFromBookshop(bookRepo repository.BookRepository, paginator *shared.Paginator) ([]aggregate.Book, error) {

	books, err := bookRepo.List(paginator.Limit(), paginator.Offset())

	if err != nil {
		return []aggregate.Book{}, err
	}

	return books, nil
}
