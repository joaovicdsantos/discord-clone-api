package model

// Channel channel model
type Channel struct {
	ID             uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	ServerID       uint        `json:"server_id" gorm:"not null"`
	Name           string      `json:"name"`
	Type           ChannelType `json:"type"`
	ChannelGroupID uint        `json:"channel_group"`
	Messages       []Message   `json:"messages" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
