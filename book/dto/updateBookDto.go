package dto

import (
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

type UpdateBookDto struct {
	bookId   valueObject.BookId
	title    valueObject.Title
	isbn     valueObject.Isbn
	year     valueObject.Year
	authorId valueObject.AuthorId
}

func NewUpdateBookDto(bookId string, title string, isbn string, year string, authorId string) (*UpdateBookDto, []error) {
	var errors []error
	NewBookId, bookIdErr := valueObject.NewBookId(bookId)
	NewTitle, titleErr := valueObject.NewTitle(title)
	NewIsbn, isbnErr := valueObject.NewIsbn(isbn)
	NewYear, yearErr := valueObject.NewYear(year)
	NewAuthorId, authorIdErr := valueObject.NewAuthorId(authorId)

	for _, err := range []error{titleErr, isbnErr, yearErr, authorIdErr, bookIdErr} {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return &UpdateBookDto{}, errors
	}

	return &UpdateBookDto{
		*NewBookId,
		*NewTitle,
		*NewIsbn,
		*NewYear,
		*NewAuthorId,
	}, errors
}

func (dto UpdateBookDto) BookId() valueObject.BookId {
	return dto.bookId
}

func (dto UpdateBookDto) Title() valueObject.Title {
	return dto.title
}

func (dto UpdateBookDto) Isbn() valueObject.Isbn {
	return dto.isbn
}

func (dto UpdateBookDto) Year() valueObject.Year {
	return dto.year
}

func (dto UpdateBookDto) AuthorId() valueObject.AuthorId {
	return dto.authorId
}
