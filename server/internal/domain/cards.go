package domain

import "github.com/gofrs/uuid"

// Cards Структура хранящихся в базе данные от карт
type Cards struct {
	Base
	UserId  uuid.UUID `gorm:"type:varchar" json:"user_id"`
	Name    string    `gorm:"type:varchar" json:"name"`
	Number  string    `gorm:"type:varchar" json:"number"`
	DateEnd string    `gorm:"type:varchar" json:"date_end"`
	CVS     string    `gorm:"type:varchar" json:"cvs"`
	Bank    string    `gorm:"type:varchar" json:"bank"`
	Meta    string    `gorm:"type:text" json:"meta"`
}
