package api

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	middleware "server/internal/app/midleware"
	"server/internal/app/services"
	"server/internal/models"
	"server/internal/tools"
)

// CardAPI структура
type CardAPI struct{}

// NewCardAPI метод возвращает указатель на структуру CardAPI со всеми ее методами
func NewCardAPI() *CardAPI {
	return &CardAPI{}
}

var cardService = services.NewCardService()

// Save сохраняет данные карты
func (ca *CardAPI) Save(c *gin.Context) {
	var body models.SaveCard

	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		tools.CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}

	value, _ := middleware.Passport().ParseTokenString(headerToken)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	cardModel, err := cardService.Save(body, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, cardModel)
}

// Get возвращает данные карты по имени
func (ca *CardAPI) Get(c *gin.Context) {
	var body models.GetCardModel

	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		tools.CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}

	value, _ := middleware.Passport().ParseTokenString(headerToken)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	cardModel, err := cardService.GetById(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	numberRes, err := tools.Base64Decode([]byte(cardModel.Number))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	cvsRes, err := tools.Base64Decode([]byte(cardModel.CVS))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	bankRes, err := tools.Base64Decode([]byte(cardModel.Bank))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	dateEndRes, err := tools.Base64Decode([]byte(cardModel.DateEnd))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("банк - %s; номер карты: %s; дейсвтует до:  %s, cvs код: %s", bankRes, numberRes, dateEndRes, cvsRes))
}

// Delete удаляет карту по имени
func (ca *CardAPI) Delete(c *gin.Context) {
	var body models.GetCardModel

	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		tools.CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}

	value, _ := middleware.Passport().ParseTokenString(headerToken)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	_, err := cardService.Delete(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, "Удаление прошло успешно")
}
