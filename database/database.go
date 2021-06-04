package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joaovicdsantos/discord-clone-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConn database connection variable
var (
	DBConn *gorm.DB
)

// InitDatabase function to start the database
func InitDatabase() {
	var (
		err      error
		DB_HOST  = os.Getenv("DB_HOST")
		DB_PORT  = os.Getenv("DB_PORT")
		DB_USER  = os.Getenv("DB_USER")
		DB_NAME  = os.Getenv("DB_NAME")
		DB_PASS  = os.Getenv("DB_PASS")
		dbLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		)
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&model.Server{})
	DBConn.AutoMigrate(&model.Channel{})
	DBConn.AutoMigrate(&model.GroupChannel{})
	DBConn.AutoMigrate(&model.User{})
	DBConn.AutoMigrate(&model.Message{})
	fmt.Println("Tables migrated")
}
