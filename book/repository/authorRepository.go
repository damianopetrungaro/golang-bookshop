package repository

import (
	"errors"
	"github.com/damianopetrungaro/golang-bookshop/book/valueObject"
)

var ImpossibleSearchAuthorError = errors.New("impossible search author")

type AuthorRepository interface {
	Exists(id valueObject.AuthorId) (bool, error)
}
