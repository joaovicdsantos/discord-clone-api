package model

// Server server model
type Server struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string         `json:"name" gorm:"unique"`
	ImageUrl      string         `json:"image_url"`
	Channels      []Channel      `json:"channels" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChannelGroups []ChannelGroup `json:"channel_groups" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Messages      []Message      `json:"messages" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
