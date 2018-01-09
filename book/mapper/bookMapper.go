package mapper

import (
	"database/sql"
	"github.com/damianopetrungaro/golang-bookshop/book/aggregate"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

type BookMapper struct {
}

type BookReadModel struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	Year     int    `json:"year"`
	AuthorId string `json:"author_id"`
}

func (mapper BookMapper) MapToBookFromRows(row *sql.Rows) *aggregate.Book {
	var id, authorId, title, isbn, year string

	if err := row.Scan(&id, &title, &isbn, &year, &authorId); err != nil {
		panic(err)
	}
	newBookId, _ := valueObject.NewBookId(id)
	newTitle, _ := valueObject.NewTitle(title)
	newIsbn, _ := valueObject.NewIsbn(isbn)
	newYear, _ := valueObject.NewYear(year)
	newAuthorId, _ := valueObject.NewAuthorId(authorId)

	return aggregate.NewBook(*newBookId, *newTitle, *newIsbn, *newYear, *newAuthorId)
}

func (mapper BookMapper) MapToPlain(book aggregate.Book) (string, string, string, int, string) {

	return book.Id().String(),
		book.Title().Value(),
		book.Isbn().Value(),
		book.Year().Value(),
		book.AuthorId().Value().String()
}

func (mapper BookMapper) MapToBookReadModel(book aggregate.Book) *BookReadModel {
	return &BookReadModel{
		book.Id().String(),
		book.Title().Value(),
		book.Isbn().Value(),
		book.Year().Value(),
		book.AuthorId().Value().String(),
	}
}
