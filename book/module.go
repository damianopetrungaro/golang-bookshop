package book

import (
	"encoding/json"
	"github.com/damianopetrungaro/golang-bookshop/book/dto"
	"github.com/damianopetrungaro/golang-bookshop/book/mapper"
	"github.com/damianopetrungaro/golang-bookshop/book/repository"
	"github.com/damianopetrungaro/golang-bookshop/book/usecase"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
	"github.com/damianopetrungaro/golang-bookshop/shared"
	"github.com/gorilla/mux"
	"net/http"
)

var container = *NewContainer()

func LoadRoutes(prefix string, router *mux.Router) {
	router.HandleFunc(prefix, list).Methods("GET")
	router.HandleFunc(prefix+"/{id}", show).Methods("GET")
	router.HandleFunc(prefix, store).Methods("POST")
	router.HandleFunc(prefix+"/{id}", update).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", remove).Methods("DELETE")
}

func list(writer http.ResponseWriter, request *http.Request) {
	container.Invoke(func(bookRepo repository.BookRepository, bookMapper *mapper.BookMapper) {
		paginator := shared.NewPaginatorFromRequest(request)

		books, err := usecase.ListBooksFromBookshop(bookRepo, paginator)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeError(&writer, err)
			return
		}

		var bookReadModelList []*mapper.BookReadModel

		for _, book := range books {
			bookReadModelList = append(bookReadModelList, bookMapper.MapToBookReadModel(book))
		}

		json.NewEncoder(writer).Encode(bookReadModelList)
	})
}

func show(writer http.ResponseWriter, request *http.Request) {
	container.Invoke(func(bookRepo repository.BookRepository, bookMapper *mapper.BookMapper) {

		id := mux.Vars(request)["id"]
		bookId, err := valueObject.NewBookId(id)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeError(&writer, err)
			return
		}

		book, err := usecase.GetBookFromBookshop(*bookId, bookRepo)

		if err != nil && err != repository.BookNotFoundError {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeError(&writer, err)
			return
		}

		if err != nil && err == repository.BookNotFoundError {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(writer).Encode(bookMapper.MapToBookReadModel(book))
	})
}

func store(writer http.ResponseWriter, request *http.Request) {

	container.Invoke(func(authorRepo repository.AuthorRepository, bookRepo repository.BookRepository, mapper *mapper.BookMapper) {
		var data = struct {
			Title    string `json:"title"`
			Isbn     string `json:"isbn"`
			Year     string `json:"year"`
			AuthorId string `json:"author_id"`
		}{}
		json.NewDecoder(request.Body).Decode(&data)
		addBookDto, errors := dto.NewAddBookDto(data.Title, data.Isbn, data.Year, data.AuthorId)

		if len(errors) > 0 {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeErrors(&writer, errors)
			return
		}

		params := usecase.AddBookToBookshopParams{Dto: *addBookDto, BookRepo: bookRepo, AuthorRepo: authorRepo}
		book, err := usecase.AddBookToBookshop(params)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeError(&writer, err)
			return
		}

		json.NewEncoder(writer).Encode(mapper.MapToBookReadModel(book))
	})
}

func update(writer http.ResponseWriter, request *http.Request) {
	container.Invoke(func(authorRepo repository.AuthorRepository, bookRepo repository.BookRepository, mapper *mapper.BookMapper) {
		var data = struct {
			BookId   string `json:"book_id"`
			Title    string `json:"title"`
			Isbn     string `json:"isbn"`
			Year     string `json:"year"`
			AuthorId string `json:"author_id"`
		}{}
		json.NewDecoder(request.Body).Decode(&data)
		data.BookId = mux.Vars(request)["id"]
		updateBookDto, errors := dto.NewUpdateBookDto(data.BookId, data.Title, data.Isbn, data.Year, data.AuthorId)

		if len(errors) > 0 {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeErrors(&writer, errors)
			return
		}

		params := usecase.UpdateBookToBookshopParams{Dto: *updateBookDto, BookRepo: bookRepo, AuthorRepo: authorRepo}
		book, err := usecase.UpdateBookInBookshop(params)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeError(&writer, err)
			return
		}

		json.NewEncoder(writer).Encode(mapper.MapToBookReadModel(book))
	})
}

func remove(writer http.ResponseWriter, request *http.Request) {
	container.Invoke(func(bookRepo repository.BookRepository, bookMapper *mapper.BookMapper) {

		id := mux.Vars(request)["id"]
		bookId, err := valueObject.NewBookId(id)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeError(&writer, err)
			return
		}

		err = usecase.DeleteBookFromBookshop(*bookId, bookRepo)

		if err != nil && err != repository.BookNotFoundError {
			writer.WriteHeader(http.StatusInternalServerError)
			encodeError(&writer, err)
			return
		}

		if err != nil && err == repository.BookNotFoundError {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	})
}

func encodeError(writer *http.ResponseWriter, err error) {
	json.NewEncoder(*writer).Encode([]string{err.Error()})
}

func encodeErrors(writer *http.ResponseWriter, errors []error) {
	var stringErrors []string

	for _, err := range errors {
		stringErrors = append(stringErrors, err.Error())
	}

	json.NewEncoder(*writer).Encode(stringErrors)
}
