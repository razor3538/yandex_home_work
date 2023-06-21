package services

import (
	"github.com/gofrs/uuid"
	repositories "server/internal/app/repository"
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
)

// TextService структура
type TextService struct{}

// NewTextService метод возвращает указатель на структуру TextService со всеми ее методами
func NewTextService() *TextService {
	return &TextService{}
}

var textRepo = repositories.NewTextRepo()

// Save сохраняет текст
func (ts *TextService) Save(textModel models.SaveText, userId string) (domain.Text, error) {
	id, err := uuid.FromString(userId)

	var text = domain.Text{
		Base:   domain.Base{},
		Text:   textModel.Text,
		Meta:   textModel.Meta,
		Name:   textModel.Name,
		UserId: id,
	}

	text.Text = string(tools.Base64Encode([]byte(text.Text)))

	result, err := textRepo.Save(text)

	if err != nil {
		return domain.Text{}, err
	}

	return result, nil
}

// GetById возвращает текст по переданному имени
func (ts *TextService) GetById(name string, userId string) (domain.Text, error) {
	result, err := textRepo.GetByKey(name, userId)
	if err != nil {
		return domain.Text{}, err
	}

	return result, nil
}

// Delete удаляет текст по переданному имени
func (ts *TextService) Delete(name string, userId string) (bool, error) {
	result, err := textRepo.Delete(name, userId)
	if err != nil {
		return false, err
	}

	return result, nil
}
