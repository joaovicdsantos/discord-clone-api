package service

import (
	"errors"
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/model"
)

type ServerService struct {
}

// FindAll find all registered users
func (s ServerService) FindAll() []model.Server {
	db := database.DBConn
	var servers []model.Server
	db.Find(&servers)
	return servers
}

// FindById find a user by id
func (s ServerService) FindById(id uint) (model.Server, error) {
	db := database.DBConn
	var server model.Server
	db.First(&server, id)
	if server.ID == 0 {
		return model.Server{}, errors.New("Not found")
	}
	return server, nil
}

// Create create a server
func (s ServerService) Create(bodyParser BodyParser) (model.Server, error) {
	db := database.DBConn
	var server model.Server
	if err := bodyParser(&server); err != nil {
		return model.Server{}, errors.New("Invalid object")
	}
	db.Save(&server)
	return server, nil
}

// Delete delete a server by id
func (s ServerService) Delete(id uint) error {
	db := database.DBConn
	db.Delete(&model.Server{}, id)
	return nil
}

// Update update a server
func (s ServerService) Update(id uint, bodyParser BodyParser) error {
	db := database.DBConn
	var server model.Server
	if err := bodyParser(&server); err != nil {
		return errors.New("Invalid object")
	}
	fmt.Println(server)
	db.Model(&model.Server{}).Where("id = ?", id).Updates(&server)
	return nil
}
