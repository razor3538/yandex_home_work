package repositories

import (
	"server/config"
	"server/internal/domain"
)

// TextRepo структура
type TextRepo struct{}

// NewTextRepo метод возвращает указатель на структуру TextRepo со всеми ее методами
func NewTextRepo() *TextRepo {
	return &TextRepo{}
}

// Save сохраняет текст
func (tr *TextRepo) Save(text domain.Text) (domain.Text, error) {
	var existingText domain.Text

	if err := config.DB.
		Table("texts as t").
		Select("t.*").
		Where("t.name = ? and t.user_id = ?", text.Name, text.UserId).
		Scan(&existingText).
		Error; err != nil {
		if err.Error() != "record not found" {
			return domain.Text{}, err
		}
	}

	if existingText.Name != "" {
		existingText.Text = text.Text

		if err := config.DB.Save(&existingText).Error; err != nil {
			return domain.Text{}, err
		}
		return text, nil
	}

	if err := config.DB.
		Create(&text).
		Error; err != nil {
		return domain.Text{}, err
	}

	return text, nil
}

// GetByKey возвращает текст по переданному имени
func (tr *TextRepo) GetByKey(value, userId string) (domain.Text, error) {
	var text domain.Text
	err := config.DB.
		Table("texts as t").
		Select("t.*").
		Where("t.name = ? and t.user_id = ?", value, userId).
		Scan(&text).Error

	return text, err
}

// Delete удаляет текст по переданному имени
func (tr *TextRepo) Delete(name, userId string) (bool, error) {
	err := config.DB.
		Table("texts as t").
		Where("t.name = ? and t.user_id = ?", name, userId).
		Delete(&domain.Text{}).Error

	if err != nil {
		return false, err
	}

	return true, err
}
