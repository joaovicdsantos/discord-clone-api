package model

// User user model
type User struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      *string   `json:"username" gorm:"not null" validate:"required,max=24,min=4,alphanumunicode"`
	Password      *string   `json:"password" gorm:"not null" validate:"required,max=24,min=8"`
	Email         *string   `json:"email" gorm:"not null;unique" validate:"required,email"`
	Status        Status    `json:"status"`
	StatusMessage string    `json:"status_message"`
	Messages      []Message `json:"mensagens" gorm:"references:ID"`
}
