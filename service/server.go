package service

import (
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/exception"
	"github.com/joaovicdsantos/discord-clone-api/model"
)

// ServerService server services class
type ServerService struct {
}

// GetAll find all registered users
func (s ServerService) GetAll() []model.Server {
	var servers []model.Server
	db := database.DBConn
	db.Find(&servers)
	return servers
}

// GetOne find a user by id
func (s ServerService) GetOne(id string) (model.Server, exception.HttpError) {

	var server model.Server

	db := database.DBConn
	if db.Preload("Channels").First(&server, id); server.ID == 0 {
		return model.Server{}, exception.HttpError{
			Err:        fmt.Errorf("server %s not found", id),
			StatusCode: 404,
		}
	}

	return server, exception.HttpError{}
}

// GetAllGroupChannels find all channels by server id
func (s ServerService) GetAllGroupChannels(id string) ([]model.ChannelGroup, exception.HttpError) {

	db := database.DBConn

	// Exists
	var server model.Server
	if db.First(&server, id); server.ID == 0 {
		return []model.ChannelGroup{}, exception.HttpError{
			Err:        fmt.Errorf("server %s not found", id),
			StatusCode: 404,
		}
	}

	var channelGroups []model.ChannelGroup
	db.Where("server_id = ?", id).Preload("Channels").Find(&channelGroups)
	return channelGroups, exception.HttpError{}
}

// Create create a server
func (s ServerService) Create(server model.Server) (model.Server, exception.HttpError) {
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
			Err:        fmt.Errorf("server %s not found", id),
			StatusCode: 404,
		}
	}

	fmt.Println(server)

	db.Delete(&model.Server{}, id)

	return exception.HttpError{}
}

// Update update a server
func (s ServerService) Update(id string, server model.Server) exception.HttpError {
	db := database.DBConn

	// Exists
	if db.First(&server, id); server.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("server %s not found", id),
			StatusCode: 404,
		}
	}

	db.Model(&model.Server{}).Where("id = ?", id).Updates(&server)

	return exception.HttpError{}
}
