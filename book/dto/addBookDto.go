package dto

import (
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

type AddBookDto struct {
	title    valueObject.Title
	isbn     valueObject.Isbn
	year     valueObject.Year
	authorId valueObject.AuthorId
}

func NewAddBookDto(title string, isbn string, year string, authorId string) (*AddBookDto, []error) {
	var errors []error
	NewTitle, titleErr := valueObject.NewTitle(title)
	NewIsbn, isbnErr := valueObject.NewIsbn(isbn)
	NewYear, yearErr := valueObject.NewYear(year)
	NewAuthorId, authorIdErr := valueObject.NewAuthorId(authorId)

	for _, err := range []error{titleErr, isbnErr, yearErr, authorIdErr} {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return &AddBookDto{}, errors
	}

	return &AddBookDto{
		*NewTitle,
		*NewIsbn,
		*NewYear,
		*NewAuthorId,
	}, errors
}

func (dto AddBookDto) Title() valueObject.Title {
	return dto.title
}

func (dto AddBookDto) Isbn() valueObject.Isbn {
	return dto.isbn
}

func (dto AddBookDto) Year() valueObject.Year {
	return dto.year
}

func (dto AddBookDto) AuthorId() valueObject.AuthorId {
	return dto.authorId
}
