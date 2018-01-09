package valueObject

import (
	"errors"
	"github.com/satori/go.uuid"
)

type BookId struct {
	value uuid.UUID
}

var BookIdInvalidFormatError = errors.New("book id must be a uuid")

func NewBookId(value string) (*BookId, error) {
	newValue, err := uuid.FromString(value)

	if err != nil {
		return &BookId{}, BookIdInvalidFormatError
	}

	return &BookId{newValue}, nil
}

func (authorId BookId) Value() uuid.UUID {
	return authorId.value
}

func (authorId BookId) String() string {
	return authorId.value.String()
}
