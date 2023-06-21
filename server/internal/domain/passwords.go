package domain

import "github.com/gofrs/uuid"

// Password Структура хранящихся в базе данных пароли и логины
type Password struct {
	Base
	UserId   uuid.UUID `gorm:"type:varchar" json:"user_id"`
	Name     string    `gorm:"type:varchar" json:"name"`
	Login    string    `gorm:"type:varchar" json:"login"`
	Password string    `gorm:"type:varchar" json:"password"`
	Meta     string    `gorm:"type:text" json:"meta"`
}
