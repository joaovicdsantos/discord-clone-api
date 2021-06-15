package service

import (
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/exception"
	"github.com/joaovicdsantos/discord-clone-api/model"
)

// MessageService message service class
type MessageService struct {
}

// GetAll find all registered messages
func (m MessageService) GetAll() []model.Message {
	var message []model.Message
	db := database.DBConn
	db.Find(&message)
	return message
}

// GetOne find a message by id
func (m MessageService) GetOne(id string) (model.Message, exception.HttpError) {

	var message model.Message

	db := database.DBConn
	if db.First(&message, id); message.ID == 0 {
		return model.Message{}, exception.HttpError{
			Err:        fmt.Errorf("message %s not found", id),
			StatusCode: 404,
		}
	}

	return message, exception.HttpError{}
}

// Create create a message
func (m MessageService) Create(message model.Message) (model.Message, exception.HttpError) {
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
			Err:        fmt.Errorf("message %s not found", id),
			StatusCode: 404,
		}
	}

	fmt.Println(message)

	db.Delete(&model.Message{}, id)

	return exception.HttpError{}
}

// Update update a message by id
func (m MessageService) Update(id string, message model.Message) exception.HttpError {
	db := database.DBConn

	// Exists
	if db.First(&message, id); message.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("message %s not found", id),
			StatusCode: 404,
		}
	}

	db.Model(&model.Message{}).Where("id = ?", id).Updates(&message)

	return exception.HttpError{}
}
