package service

import "github.com/joaovicdsantos/discord-clone-api/exception"

// BodyParser a fiber BodyParser abstraction
type BodyParser func(out interface{}) error

// Service service interface example
type Service interface {
	FindAll() []interface{}

	FindById(stringId string) (interface{}, exception.HttpError)

	Create(bodyParser BodyParser) exception.HttpError

	Delete(stringId string) exception.HttpError

	Update(stringId string, bodyParser BodyParser) (interface{}, exception.HttpError)
}
