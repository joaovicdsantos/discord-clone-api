package model

import "gorm.io/gorm"

type ChannelType string

const (
	Text  ChannelType = "text"
	Voice ChannelType = "voice"
)

type Server struct {
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string    `json:"name" gorm:"unique"`
	Channels []Channel `json:"channels" gorm:"references:ID"`
}

type Channel struct {
	ID             uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	ServerID       uint         `json:"server_id"`
	Name           string       `json:"name"`
	Type           ChannelType  `json:"type"`
	GroupChannelID uint         `json:"group_channel"`
	Group          GroupChannel `json:"group" gorm:"foreignKey:GroupChannelID;references:ID"`
}

type GroupChannel struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID    uint      `json:"group_id"`
	Name       string    `json:"name"`
	ChannelsID uint      `json:"channels_id"`
	Channels   []Channel `json:"channels" gorm:"references:ID"`
}

type User struct {
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Messages []Message `json:"mensagens" gorm:"references:ID"`
}

type Message struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}
