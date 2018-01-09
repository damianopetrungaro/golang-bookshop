package usecase

import (
	"errors"
	"github.com/damianopetrungaro/golang-bookshop/book/aggregate"
	"github.com/damianopetrungaro/golang-bookshop/book/dto"
	"github.com/damianopetrungaro/golang-bookshop/book/repository"
)

type AddBookToBookshopParams struct {
	Dto        dto.AddBookDto
	BookRepo   repository.BookRepository
	AuthorRepo repository.AuthorRepository
}

var AuthorNotFoundError = errors.New("author not found")

func AddBookToBookshop(params AddBookToBookshopParams) (aggregate.Book, error) {

	exists, err := params.AuthorRepo.Exists(params.Dto.AuthorId())

	if err != nil {
		return aggregate.Book{}, err
	}

	if !exists {
		return aggregate.Book{}, AuthorNotFoundError
	}

	book := *aggregate.NewBook(
		params.BookRepo.NextId(),
		params.Dto.Title(),
		params.Dto.Isbn(),
		params.Dto.Year(),
		params.Dto.AuthorId(),
	)

	if err = params.BookRepo.Add(book); err != nil {
		return aggregate.Book{}, err
	}

	return book, nil
}
