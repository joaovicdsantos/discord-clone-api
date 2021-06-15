package service

import (
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/exception"
	"github.com/joaovicdsantos/discord-clone-api/model"
)

// ChannelGroupService channel group service class
type ChannelGroupService struct {
}

// GetAll find all channel groups
func (c ChannelGroupService) GetAll() []model.ChannelGroup {
	var channelGroups []model.ChannelGroup
	db := database.DBConn
	db.Find(&channelGroups)
	return channelGroups
}

// GetOne find a channel group by id
func (c ChannelGroupService) GetOne(id string) (model.ChannelGroup, exception.HttpError) {

	var channelGroup model.ChannelGroup

	db := database.DBConn
	if db.First(&channelGroup, id); channelGroup.ID == 0 {
		return model.ChannelGroup{}, exception.HttpError{
			Err:        fmt.Errorf("channel group %s not found", id),
			StatusCode: 404,
		}
	}

	return channelGroup, exception.HttpError{}
}

// Create create a channel group
func (c ChannelGroupService) Create(channelGroup model.ChannelGroup) (model.ChannelGroup, exception.HttpError) {
	db := database.DBConn
	db.Save(&channelGroup)

	return channelGroup, exception.HttpError{}
}

// Delete delete a channel group by id
func (c ChannelGroupService) Delete(id string) exception.HttpError {
	db := database.DBConn
	// Exists
	var channelGroup model.ChannelGroup
	if db.First(&channelGroup, id); channelGroup.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("channel group %s not found", id),
			StatusCode: 404,
		}
	}

	fmt.Println(channelGroup)

	db.Delete(&model.ChannelGroup{}, id)

	return exception.HttpError{}
}

// Update update a channel group by id
func (c ChannelGroupService) Update(id string, channelGroup model.ChannelGroup) exception.HttpError {
	db := database.DBConn

	// Exists
	if db.First(&channelGroup, id); channelGroup.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("channel group %s not found", id),
			StatusCode: 404,
		}
	}

	db.Model(&model.ChannelGroup{}).Where("id = ?", id).Updates(&channelGroup)

	return exception.HttpError{}
}
