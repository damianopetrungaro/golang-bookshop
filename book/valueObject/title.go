package valueObject

import "errors"

type Title struct {
	value string
}

var TitleInvalidLengthError = errors.New("title length must be between 10 and 50")

func NewTitle(value string) (*Title, error) {

	if len(value) < 10 || len(value) > 50 {
		return &Title{}, TitleInvalidLengthError
	}

	return &Title{value}, nil
}

func (title Title) Value() string {
	return title.value
}
