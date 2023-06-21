package domain

// User Структура хранящихся в базе данных пользователей
type User struct {
	Base
	Login    string `gorm:"type:varchar" json:"login"`
	Password string `gorm:"type:varchar" json:"password"`
}
