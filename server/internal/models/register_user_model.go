package models

// SaveUserRequest модель для сохранения пользователя
type SaveUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
