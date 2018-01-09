package usecase

import (
	"github.com/damianopetrungaro/golang-bookshop/book/repository"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

func DeleteBookFromBookshop(bookId valueObject.BookId, bookRepo repository.BookRepository) error {

	if _, err := bookRepo.Get(bookId); err != nil {
		return err
	}

	if err := bookRepo.Delete(bookId); err != nil {
		return err
	}

	return nil
}
