package usecase

import (
	"github.com/damianopetrungaro/golang-bookshop/book/aggregate"
	"github.com/damianopetrungaro/golang-bookshop/book/dto"
	"github.com/damianopetrungaro/golang-bookshop/book/repository"
)

type UpdateBookToBookshopParams struct {
	Dto        dto.UpdateBookDto
	BookRepo   repository.BookRepository
	AuthorRepo repository.AuthorRepository
}

func UpdateBookInBookshop(params UpdateBookToBookshopParams) (aggregate.Book, error) {

	exists, err := params.AuthorRepo.Exists(params.Dto.AuthorId())

	if err != nil {
		return aggregate.Book{}, err
	}

	if !exists {
		return aggregate.Book{}, AuthorNotFoundError
	}

	book, err := params.BookRepo.Get(params.Dto.BookId())

	book = book.Update(
		params.Dto.Title(),
		params.Dto.Isbn(),
		params.Dto.Year(),
		params.Dto.AuthorId(),
	)

	if err = params.BookRepo.Update(book); err != nil {
		return aggregate.Book{}, err
	}

	return book, nil
}
