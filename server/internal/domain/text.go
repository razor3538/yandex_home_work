package domain

import "github.com/gofrs/uuid"

// Text Структура хранящихся в базе данных пароли и логины
type Text struct {
	Base
	UserId uuid.UUID `gorm:"type:varchar" json:"user_id"`
	Name   string    `gorm:"type:varchar" json:"name"`
	Text   string    `gorm:"type:varchar" json:"text"`
	Meta   string    `gorm:"type:text" json:"meta"`
}
