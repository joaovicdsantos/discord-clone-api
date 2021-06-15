package utils

import (
	"errors"

	"github.com/joaovicdsantos/discord-clone-api/exception"
)

// BodyParser a fiber BodyParser abstraction
type BodyParser func(out interface{}) error

// ToModel returns an object formed from body
func ToModel(model interface{}, bodyParser BodyParser) exception.HttpError {
	if err := bodyParser(&model); err != nil {
		return exception.HttpError{
			Err:        errors.New("invalid object"),
			StatusCode: 400,
		}
	}
	return exception.HttpError{}
}
