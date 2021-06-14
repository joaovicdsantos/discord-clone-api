package service

import (
	"errors"
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/exception"
	"github.com/joaovicdsantos/discord-clone-api/model"
)

// ServerService server services class
type ServerService struct {
}

// FindAll find all registered users
func (s ServerService) FindAll() []model.Server {
	var servers []model.Server
	db := database.DBConn
	db.Find(&servers)
	return servers
}

// FindById find a user by id
func (s ServerService) FindById(id string) (model.Server, exception.HttpError) {

	var server model.Server

	db := database.DBConn
	if db.Preload("Channels").First(&server, id); server.ID == 0 {
		return model.Server{}, exception.HttpError{
			Err:        fmt.Errorf("Server %s not found.", id),
			StatusCode: 404,
		}
	}

	return server, exception.HttpError{}
}

// FindAllGroupChannels find all channels by server id
func (s ServerService) FindAllGroupChannels(id string) ([]model.ChannelGroup, exception.HttpError) {

	db := database.DBConn

	// Exists
	var server model.Server
	if db.First(&server, id); server.ID == 0 {
		return []model.ChannelGroup{}, exception.HttpError{
			Err:        fmt.Errorf("Server %s not found.", id),
			StatusCode: 404,
		}
	}

	var channelGroups []model.ChannelGroup
	db.Where("server_id = ?", id).Preload("Channels").Find(&channelGroups)
	return channelGroups, exception.HttpError{}
}

// Create create a server
func (s ServerService) Create(bodyParser BodyParser) (model.Server, exception.HttpError) {
	var server model.Server
	if err := bodyParser(&server); err != nil {
		return model.Server{}, exception.HttpError{
			Err:        errors.New("Invalid object."),
			StatusCode: 400,
		}
	}

	db := database.DBConn
	db.Save(&server)

	fmt.Print(server.ID)

	var channelGroup model.ChannelGroup
	channelGroup.Name = "default"
	channelGroup.ServerID = server.ID
	db.Save(&channelGroup)

	return server, exception.HttpError{}
}

// Delete delete a server by id
func (s ServerService) Delete(id string) exception.HttpError {
	db := database.DBConn
	// Exists
	var server model.Server
	if db.First(&server, id); server.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("Server %s not found.", id),
			StatusCode: 404,
		}
	}

	fmt.Println(server)

	db.Delete(&model.Server{}, id)

	return exception.HttpError{}
}

// Update update a server
func (s ServerService) Update(id string, bodyParser BodyParser) exception.HttpError {
	var server model.Server
	if err := bodyParser(&server); err != nil {
		return exception.HttpError{
			Err:        errors.New("Invalid object"),
			StatusCode: 400,
		}
	}

	db := database.DBConn

	// Exists
	if db.First(&server, id); server.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("Server %s not found.", id),
			StatusCode: 404,
		}
	}

	db.Model(&model.Server{}).Where("id = ?", id).Updates(&server)

	return exception.HttpError{}
}
