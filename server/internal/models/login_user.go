package models

// LoginUserModel order model for login user
type LoginUserModel struct {
	Username string `json:"login" example:"admin@mail.ru" binding:"required"`
	Password string `json:"password" example:"123" binding:"required"`
} //@name LoginUserModel
