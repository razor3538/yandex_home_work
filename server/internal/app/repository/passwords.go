package repositories

import (
	"server/config"
	"server/internal/domain"
)

// PasswordRepo структура
type PasswordRepo struct{}

// NewPasswordRepo метод возвращает указатель на структуру PasswordRepo со всеми ее методами
func NewPasswordRepo() *PasswordRepo {
	return &PasswordRepo{}
}

// Save сохраняет пару логин + пароль
func (pr *PasswordRepo) Save(pass domain.Password) (domain.Password, error) {
	var existingPass domain.Password

	if err := config.DB.
		Table("passwords as p").
		Select("p.*").
		Where("p.name = ? and p.user_id = ?", pass.Name, pass.UserId).
		Scan(&existingPass).
		Error; err != nil {
		if err.Error() != "record not found" {
			return domain.Password{}, err
		}
	}

	if existingPass.Name != "" {
		existingPass.Login = pass.Login
		existingPass.Password = pass.Password

		if err := config.DB.Save(&existingPass).Error; err != nil {
			return domain.Password{}, err
		}
		return pass, nil
	}

	if err := config.DB.
		Create(&pass).
		Error; err != nil {
		return domain.Password{}, err
	}

	return pass, nil
}

// GetByKey возвращает пару логин + пароль по переданному имени
func (pr *PasswordRepo) GetByKey(value, userId string) (domain.Password, error) {
	var pass domain.Password
	err := config.DB.
		Table("passwords as p").
		Select("p.*").
		Where("p.name = ? and p.user_id = ?", value, userId).
		Scan(&pass).Error

	return pass, err
}

// Delete удаляет пару логин + пароль по переданному имени
func (pr *PasswordRepo) Delete(name, userId string) (bool, error) {
	err := config.DB.
		Table("passwords as p").
		Where("p.name = ? and p.user_id = ?", name, userId).
		Delete(&domain.Password{}).Error

	if err != nil {
		return false, err
	}

	return true, err
}
