package repositories

import (
	"errors"
	"server/config"
	"server/internal/domain"
)

// UserRepo структура
type UserRepo struct{}

// NewUserRepo метод возвращает указатель на структуру UserRepo со всеми ее методами
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// Save сохраняет пользователя
func (ur *UserRepo) Save(user domain.User) (domain.User, error) {
	var existingUser domain.User

	if err := config.DB.
		Table("users as u").
		Select("u.*").
		Where("u.login = ?", user.Login).
		Scan(&existingUser).
		Error; err != nil {
		if err.Error() != "record not found" {
			return domain.User{}, err
		}
	}

	if existingUser.Login != "" {
		return existingUser, errors.New("пользователь уже существует")
	}

	if err := config.DB.
		Create(&user).
		Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// GetByKey возвращает пользователя по ключу
func (ur *UserRepo) GetByKey(key, value string) (domain.User, error) {
	var user domain.User
	err := config.DB.
		Unscoped().
		Where(key+" = ?", value).
		First(&user).Error

	return user, err
}

// Get возвращает пользователя
func (ur *UserRepo) Get() ([]domain.User, error) {
	var user []domain.User

	err := config.DB.
		Table("users as u").
		Select("u.*").
		Scan(&user).
		Error

	return user, err
}
