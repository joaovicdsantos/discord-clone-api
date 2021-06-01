package database

import (
	"fmt"
	"os"

	"github.com/joaovicdsantos/discord-clone-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var (
		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
		DB_USER = os.Getenv("DB_USER")
		DB_NAME = os.Getenv("DB_NAME")
		DB_PASS = os.Getenv("DB_PASS")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	DBConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&model.Server{})
	DBConn.AutoMigrate(&model.Channel{})
	DBConn.AutoMigrate(&model.GroupChannel{})
	DBConn.AutoMigrate(&model.User{})
	DBConn.AutoMigrate(&model.Message{})
}
