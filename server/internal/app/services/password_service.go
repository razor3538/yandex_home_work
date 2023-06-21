package services

import (
	"github.com/gofrs/uuid"
	repositories "server/internal/app/repository"
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
)

// PasswordService структура
type PasswordService struct{}

// NewPasswordService метод возвращает указатель на структуру PasswordService со всеми ее методами
func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

var passRepo = repositories.NewPasswordRepo()

// Save сохраняет пару логин + пароль
func (ps *PasswordService) Save(passModel models.SavePassword, userId string) (domain.Password, error) {
	id, err := uuid.FromString(userId)

	var pass = domain.Password{
		Base:     domain.Base{},
		Login:    passModel.Login,
		Password: passModel.Password,
		Meta:     passModel.Meta,
		Name:     passModel.Name,
		UserId:   id,
	}

	pass.Password = string(tools.Base64Encode([]byte(pass.Password)))
	pass.Login = string(tools.Base64Encode([]byte(pass.Login)))

	result, err := passRepo.Save(pass)

	if err != nil {
		return domain.Password{}, err
	}

	return result, nil
}

// GetById возвращает пару логин + пароль по переданному имени
func (ps *PasswordService) GetById(name string, userId string) (domain.Password, error) {
	result, err := passRepo.GetByKey(name, userId)
	if err != nil {
		return domain.Password{}, err
	}

	return result, nil
}

// Delete удаляет пару логин + пароль по переданному имени
func (ps *PasswordService) Delete(name string, userId string) (bool, error) {
	result, err := passRepo.Delete(name, userId)
	if err != nil {
		return false, err
	}

	return result, nil
}
