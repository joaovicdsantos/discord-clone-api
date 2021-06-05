package service

import (
	"errors"
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/exception"
	"github.com/joaovicdsantos/discord-clone-api/model"
)

// MessageService message service class
type MessageService struct {
}

// FindAll find all registered messages
func (m MessageService) FindAll() []model.Message {
	var message []model.Message
	db := database.DBConn
	db.Find(&message)
	return message
}

// FindById find a message by id
func (m MessageService) FindById(id string) (model.Message, exception.HttpError) {

	var message model.Message

	db := database.DBConn
	if db.First(&message, id); message.ID == 0 {
		return model.Message{}, exception.HttpError{
			Err:        fmt.Errorf("Message %s not found.", id),
			StatusCode: 404,
		}
	}

	return message, exception.HttpError{}
}

// Create create a message
func (m MessageService) Create(bodyParser BodyParser) (model.Message, exception.HttpError) {
	var message model.Message
	if err := bodyParser(&message); err != nil {
		return model.Message{}, exception.HttpError{
			Err:        errors.New("Invalid object."),
			StatusCode: 400,
		}
	}

	db := database.DBConn
	db.Save(&message)

	return message, exception.HttpError{}
}

// Delete delete a message by id
func (m MessageService) Delete(id string) exception.HttpError {
	db := database.DBConn
	// Exists
	var message model.Message
	if db.First(&message, id); message.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("Message %s not found.", id),
			StatusCode: 404,
		}
	}

	fmt.Println(message)

	db.Delete(&model.Message{}, id)

	return exception.HttpError{}
}

// Update update a message by id
func (m MessageService) Update(id string, bodyParser BodyParser) exception.HttpError {
	var message model.Message
	if err := bodyParser(&message); err != nil {
		return exception.HttpError{
			Err:        errors.New("Invalid object"),
			StatusCode: 400,
		}
	}

	db := database.DBConn

	// Exists
	if db.First(&message, id); message.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("Message %s not found.", id),
			StatusCode: 404,
		}
	}

	db.Model(&model.Message{}).Where("id = ?", id).Updates(&message)

	return exception.HttpError{}
}
