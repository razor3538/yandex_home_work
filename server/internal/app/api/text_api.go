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

// TextAPI структура
type TextAPI struct{}

// NewTextAPI метод возвращает указатель на структуру TextAPI со всеми ее методами
func NewTextAPI() *TextAPI {
	return &TextAPI{}
}

var textService = services.NewTextService()

// Save сохраняет текст
func (ta *TextAPI) Save(c *gin.Context) {
	var body models.SaveText

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

	textModel, err := textService.Save(body, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, textModel)
}

// Get возвращает текст по переданному имени
func (ta *TextAPI) Get(c *gin.Context) {
	var body models.GetText

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

	textModel, err := textService.GetById(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	textRes, err := tools.Base64Decode([]byte(textModel.Text))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("%s: %s", textModel.Name, textRes))
}

// Delete удаляет текст по переданному имени
func (ta *TextAPI) Delete(c *gin.Context) {
	var body models.GetText

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

	_, err := textService.Delete(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, "Удаление прошло успешно")
}
