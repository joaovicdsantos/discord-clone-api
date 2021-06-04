package model

import "gorm.io/gorm"

// ChannelType types of channels
type ChannelType string

// Text text channel
// VOice voice channel
const (
	Text  ChannelType = "text"
	Voice ChannelType = "voice"
)

// Server server model
type Server struct {
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string    `json:"name" gorm:"unique"`
	Channels []Channel `json:"channels" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Channel channel model
type Channel struct {
	ID             uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	ServerID       uint        `json:"server_id" gorm:"not null"`
	Name           string      `json:"name"`
	Type           ChannelType `json:"type"`
	ChannelGroupID uint        `json:"channel_group"`
}

// ChannelGroup channel group model
type ChannelGroup struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID    uint      `json:"group_id"`
	Name       string    `json:"name"`
	ChannelsID uint      `json:"channels_id"`
	Channels   []Channel `json:"channels" gorm:"references:ID"`
}

// User user model
type User struct {
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Messages []Message `json:"mensagens" gorm:"references:ID"`
}

// Message message model
type Message struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}
