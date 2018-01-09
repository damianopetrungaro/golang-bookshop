package aggregate

import (
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

type Book struct {
	id       valueObject.BookId
	title    valueObject.Title
	isbn     valueObject.Isbn
	year     valueObject.Year
	authorId valueObject.AuthorId
}

func NewBook(
	id valueObject.BookId,
	title valueObject.Title,
	isbn valueObject.Isbn,
	year valueObject.Year,
	author valueObject.AuthorId,
) *Book {
	return &Book{id, title, isbn, year, author}
}

func (book Book) Update(title valueObject.Title, isbn valueObject.Isbn, year valueObject.Year, authorId valueObject.AuthorId) Book {
	return Book{book.id, title, isbn, year, authorId}
}

func (book Book) Id() valueObject.BookId {
	return book.id
}

func (book Book) Title() valueObject.Title {
	return book.title
}

func (book Book) Isbn() valueObject.Isbn {
	return book.isbn
}

func (book Book) Year() valueObject.Year {
	return book.year
}

func (book Book) AuthorId() valueObject.AuthorId {
	return book.authorId
}
