package services

import (
	"github.com/gofrs/uuid"
	repositories "server/internal/app/repository"
	"server/internal/domain"
	"server/internal/models"
	"server/internal/tools"
)

// CardService структура
type CardService struct{}

// NewCardService метод возвращает указатель на структуру CardService со всеми ее методами
func NewCardService() *CardService {
	return &CardService{}
}

var cardRepo = repositories.NewCardRepo()

// Save сохраняет данные карты
func (cs *CardService) Save(cardModel models.SaveCard, userId string) (domain.Cards, error) {
	id, err := uuid.FromString(userId)

	var card = domain.Cards{
		Base:    domain.Base{},
		UserId:  id,
		Name:    cardModel.Name,
		Number:  cardModel.Number,
		DateEnd: cardModel.DateEnd,
		CVS:     cardModel.CVS,
		Bank:    cardModel.Bank,
		Meta:    cardModel.Meta,
	}

	card.Number = string(tools.Base64Encode([]byte(card.Number)))
	card.CVS = string(tools.Base64Encode([]byte(card.CVS)))
	card.DateEnd = string(tools.Base64Encode([]byte(card.DateEnd)))
	card.Bank = string(tools.Base64Encode([]byte(card.Bank)))

	result, err := cardRepo.Save(card)

	if err != nil {
		return domain.Cards{}, err
	}

	return result, nil
}

// GetById возвращает данные карты по имени
func (cs *CardService) GetById(name string, userId string) (domain.Cards, error) {
	result, err := cardRepo.GetByKey(name, userId)
	if err != nil {
		return domain.Cards{}, err
	}

	return result, nil
}

// Delete удаляет карту по имени
func (cs *CardService) Delete(name string, userId string) (bool, error) {
	result, err := cardRepo.Delete(name, userId)
	if err != nil {
		return false, err
	}

	return result, nil
}
