package services

import (
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
)

// UserService структура
type UserService struct{}

// NewUserService метод возвращает указатель на структуру UserService со всеми ее методами
func NewUserService() *UserService {
	return &UserService{}
}

// Save сохраняет пользователя
func (us *UserService) Save(userModel models.SaveUserRequest) (domain.User, error) {
	var user = domain.User{
		Base:     domain.Base{},
		Login:    userModel.Login,
		Password: userModel.Password,
	}

	user.Password = tools.HashString(user.Password)

	result, err := userRepo.Save(user)

	if err != nil {
		return domain.User{}, err
	}

	return result, nil
}

// Get возвращает пользователя
func (us *UserService) Get() ([]domain.User, error) {
	users, err := userRepo.Get()

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}
