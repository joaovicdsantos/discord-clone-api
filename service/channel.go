package service

import (
	"errors"
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/exception"
	"github.com/joaovicdsantos/discord-clone-api/model"
)

// ChannelService channel services class
type ChannelService struct {
}

// FindAll find all registered users
func (c ChannelService) FindAll() []model.Channel {
	var channels []model.Channel
	db := database.DBConn
	db.Find(&channels)
	return channels
}

// FindById find a user by id
func (c ChannelService) FindById(id string) (model.Channel, exception.HttpError) {

	var channel model.Channel

	db := database.DBConn
	if db.Preload("Messages").Where("id = ?", id).Find(&channel); channel.ID == 0 {
		return model.Channel{}, exception.HttpError{
			Err:        fmt.Errorf("Channel %s not found.", id),
			StatusCode: 404,
		}
	}

	return channel, exception.HttpError{}
}

// Create create a server
func (c ChannelService) Create(bodyParser BodyParser) (model.Channel, exception.HttpError) {
	var channel model.Channel
	if err := bodyParser(&channel); err != nil {
		return model.Channel{}, exception.HttpError{
			Err:        errors.New("Invalid object."),
			StatusCode: 400,
		}
	}

	// Verify serverID
	if channel.ServerID == 0 {
		return model.Channel{}, exception.HttpError{
			Err:        errors.New("Server ID is required"),
			StatusCode: 400,
		}
	}

	db := database.DBConn

	// This server id exists?
	var serverService ServerService
	_, err := serverService.FindById(fmt.Sprint(channel.ServerID))
	if err.Err != nil {
		return model.Channel{}, err
	}

	// This channel group exists?
	if channel.ChannelGroupID != 0 {
		var groupChannelService ChannelGroupService
		_, err := groupChannelService.FindById(fmt.Sprint(channel.ChannelGroupID))
		if err.Err != nil {
			return model.Channel{}, err
		}
	} else {
		channel.ChannelGroupID = 1
	}

	db.Save(&channel)

	return channel, exception.HttpError{}
}

// Delete delete a server by id
func (c ChannelService) Delete(id string) exception.HttpError {
	db := database.DBConn
	// Exists
	var channel model.Channel
	if db.First(&channel, id); channel.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("Channel %s not found.", id),
			StatusCode: 404,
		}
	}

	fmt.Println(channel)

	db.Delete(&model.Channel{}, id)

	return exception.HttpError{}
}

// Update update a server
func (c ChannelService) Update(id string, bodyParser BodyParser) exception.HttpError {
	var channel model.Channel
	if err := bodyParser(&channel); err != nil {
		return exception.HttpError{
			Err:        errors.New("Invalid object"),
			StatusCode: 400,
		}
	}

	db := database.DBConn

	// Exists
	if db.First(&channel, id); channel.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("Channel %s not found.", id),
			StatusCode: 404,
		}
	}

	db.Model(&model.Channel{}).Where("id = ?", id).Updates(&channel)

	return exception.HttpError{}
}
