package repositories

import (
	"server/config"
	"server/internal/domain"
)

// CardRepo структура
type CardRepo struct{}

// NewCardRepo метод возвращает указатель на структуру CardRepo со всеми ее методами
func NewCardRepo() *CardRepo {
	return &CardRepo{}
}

// Save сохраняет данные карты
func (cr *CardRepo) Save(card domain.Cards) (domain.Cards, error) {
	var existingCard domain.Cards

	if err := config.DB.
		Table("cards as c").
		Select("c.*").
		Where("c.name = ? and c.user_id = ?", card.Name, card.UserId).
		Scan(&existingCard).
		Error; err != nil {
		if err.Error() != "record not found" {
			return domain.Cards{}, err
		}
	}

	if existingCard.Name != "" {
		existingCard.CVS = card.CVS
		existingCard.Bank = card.Bank
		existingCard.Number = card.Number
		existingCard.DateEnd = card.DateEnd

		if err := config.DB.Save(&existingCard).Error; err != nil {
			return domain.Cards{}, err
		}
		return card, nil
	}

	if err := config.DB.
		Create(&card).
		Error; err != nil {
		return domain.Cards{}, err
	}

	return card, nil
}

// GetByKey возвращает данные карты по имени
func (cr *CardRepo) GetByKey(value, userId string) (domain.Cards, error) {
	var card domain.Cards
	err := config.DB.
		Table("cards as c").
		Select("c.*").
		Where("c.name = ? and c.user_id = ?", value, userId).
		Scan(&card).Error

	return card, err
}

// Delete удаляет карту по имени
func (cr *CardRepo) Delete(name, userId string) (bool, error) {
	err := config.DB.
		Table("cards as c").
		Where("c.name = ? and c.user_id = ?", name, userId).
		Delete(&domain.Cards{}).Error

	if err != nil {
		return false, err
	}

	return true, err
}
