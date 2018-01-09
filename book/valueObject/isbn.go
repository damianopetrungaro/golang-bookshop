package valueObject

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
)

type Isbn struct {
	value string
}

var IsbnInvalidLengthError = errors.New("isbn must be length 13")

func NewIsbn(value string) (*Isbn, error) {

	match, err := regexp.MatchString(`^\d{13}$`, value)

	fmt.Println(match, err, value)

	if err != nil {
		log.Error(err)
	}

	if match == false {
		return &Isbn{}, IsbnInvalidLengthError
	}

	return &Isbn{value}, nil
}

func (isbn Isbn) Value() string {
	return isbn.value
}
