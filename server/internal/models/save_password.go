package models

// SavePassword модель для сохранения пары логин + пароль
type SavePassword struct {
	Login    string `json:"login"`
	Meta     string `json:"meta"`
	Password string `json:"password"`
	Name     string `json:"name_pair"`
}
