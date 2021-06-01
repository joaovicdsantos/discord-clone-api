package service

type BodyParser func(out interface{}) error

type Service interface {
	FindAll() interface{}

	FindById(id uint) (interface{}, error)

	Create(bodyParser BodyParser) error

	Delete(id uint) error

	Update(id uint, bodyParser BodyParser) (interface{}, error)
}
