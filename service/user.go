package service

import (
	"errors"
	"fmt"

	"github.com/joaovicdsantos/discord-clone-api/database"
	"github.com/joaovicdsantos/discord-clone-api/exception"
	"github.com/joaovicdsantos/discord-clone-api/jwt"
	"github.com/joaovicdsantos/discord-clone-api/model"
	"github.com/joaovicdsantos/discord-clone-api/validator"
	"golang.org/x/crypto/bcrypt"
)

// UserService channel group service class
type UserService struct {
}

// FindAll find all registered channel groups
func (u UserService) FindAll() []model.User {
	var users []model.User
	db := database.DBConn
	db.Find(&users)
	return users
}

// FindById find a channel group by id
func (u UserService) FindById(id string) (model.User, exception.HttpError) {

	var user model.User

	db := database.DBConn
	if db.First(&user, id); user.ID == 0 {
		return model.User{}, exception.HttpError{
			Err:        fmt.Errorf("User %s not found.", id),
			StatusCode: 404,
		}
	}

	return user, exception.HttpError{}
}

// Login log in with a user
func (u UserService) Login(bodyParser BodyParser) (string, exception.HttpError) {
	var user model.User
	if err := bodyParser(&user); err != nil {
		return "", exception.HttpError{
			Err:        errors.New("Invalid object"),
			StatusCode: 400,
		}
	}

	db := database.DBConn

	// Verifications
	var registeredUser model.User
	db.Where("username = ?", user.Username).Find(&registeredUser)
	if registeredUser.ID == 0 {
		return "", exception.HttpError{
			Err:        errors.New("Username not found"),
			StatusCode: 401,
		}
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(*registeredUser.Password), []byte(*user.Password)); err != nil {
		return "", exception.HttpError{
			Err:        errors.New("Invalid password"),
			StatusCode: 401,
		}
	}

	token, err := jwt.GenerateToken(map[string]string{
		"username": *user.Username,
	})
	if err != nil {
		return "", exception.HttpError{
			Err:        err,
			StatusCode: 500,
		}
	}

	return token, exception.HttpError{}

}

// Create create a channel group
func (u UserService) Create(bodyParser BodyParser) (model.User, exception.HttpError) {
	var user model.User
	if err := bodyParser(&user); err != nil {
		return model.User{}, exception.HttpError{
			Err:        errors.New("Internal error"),
			StatusCode: 500,
		}
	}

	// Validation
	errs := validator.Validation(user)
	if errs != nil {
		var errString string
		for _, err := range errs {
			errString = fmt.Sprintf("%s\n%s failed in %s tag", errString, err.Field(), err.Tag())
		}
		return model.User{}, exception.HttpError{
			Err:        errors.New(errString),
			StatusCode: 400,
		}
	}

	// Hashing password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, exception.HttpError{
			Err:        errors.New("Internal Error"),
			StatusCode: 500,
		}
	}
	encryptedPasswordString := string(encryptedPassword)
	user.Password = &encryptedPasswordString

	db := database.DBConn
	db.Save(&user)

	return user, exception.HttpError{}
}

// Delete delete a channel group by id
func (u UserService) Delete(id string) exception.HttpError {
	db := database.DBConn
	// Exists
	var user model.User
	if db.First(&user, id); user.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("User %s not found.", id),
			StatusCode: 404,
		}
	}

	fmt.Println(user)

	db.Delete(&model.User{}, id)

	return exception.HttpError{}
}

// Update update a channel group by id
func (u UserService) Update(id string, bodyParser BodyParser) exception.HttpError {
	var user model.User
	if err := bodyParser(&user); err != nil {
		return exception.HttpError{
			Err:        errors.New("Invalid object"),
			StatusCode: 400,
		}
	}

	db := database.DBConn

	// Exists
	if db.First(&user, id); user.ID == 0 {
		return exception.HttpError{
			Err:        fmt.Errorf("User %s not found.", id),
			StatusCode: 404,
		}
	}

	db.Model(&model.User{}).Where("id = ?", id).Updates(&user)

	return exception.HttpError{}
}
