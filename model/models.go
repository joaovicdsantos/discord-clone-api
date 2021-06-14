package model

import (
	"gorm.io/gorm"
)

// ChannelType types of channels
type ChannelType string

// Status type of users status
type Status string

// Text text channel
// Voice voice channel
const (
	Text  ChannelType = "text"
	Voice ChannelType = "voice"
)

// Online online status
// Idle idle status
// Busy busy status
// Invisible invisible status
const (
	Online    Status = "Online"
	Idle      Status = "Idle"
	Busy      Status = "Do not disturb"
	Invisible Status = "Invisible"
)

// ChannelGroup channel group model
type ChannelGroup struct {
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string    `json:"name"`
	ServerID uint      `json:"server_id"`
	Channels []Channel `json:"channels" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Message message model
type Message struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	ServerID  uint   `json:"server_id"`
	ChannelID uint   `json:"channel_id"`
}
