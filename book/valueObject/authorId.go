package valueObject

import (
	"errors"
	"github.com/satori/go.uuid"
)

type AuthorId struct {
	value uuid.UUID
}

var AuthorIdInvalidFormatError = errors.New("author id must be a uuid")

func NewAuthorId(value string) (*AuthorId, error) {
	newValue, err := uuid.FromString(value)

	if err != nil {
		return &AuthorId{}, AuthorIdInvalidFormatError
	}

	return &AuthorId{newValue}, nil
}

func (authorId AuthorId) Value() uuid.UUID {
	return authorId.value
}

func (authorId AuthorId) String() string {
	return authorId.value.String()
}
