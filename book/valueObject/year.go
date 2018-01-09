package valueObject

import (
	"errors"
	"strconv"
)

type Year struct {
	value int
}

var YearInvalidFormatError = errors.New("year must be a 4 digit value")

func NewYear(value string) (*Year, error) {

	newValue, err := strconv.Atoi(value)

	if err != nil {
		return &Year{}, YearInvalidFormatError
	}

	return &Year{newValue}, nil
}

func (year Year) Value() int {
	return year.value
}

func (year Year) String() string {
	return string(year.value)
}
