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

// PasswordAPI структура
type PasswordAPI struct{}

// NewPasswordAPI метод возвращает указатель на структуру PasswordAPI со всеми ее методами
func NewPasswordAPI() *PasswordAPI {
	return &PasswordAPI{}
}

var passService = services.NewPasswordService()

// Save сохраняет пару логин + пароль
func (pa *PasswordAPI) Save(c *gin.Context) {
	var body models.SavePassword

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

	passModel, err := passService.Save(body, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, passModel)
}

// Get возвращает пару логин + пароль по переданному имени
func (pa *PasswordAPI) Get(c *gin.Context) {
	var body models.GetPassword

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

	passModel, err := passService.GetById(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	loginRes, err := tools.Base64Decode([]byte(passModel.Login))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	passRes, err := tools.Base64Decode([]byte(passModel.Password))
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("%s - %s", loginRes, passRes))
}

// Delete удаляет пару логин + пароль по переданному имени
func (pa *PasswordAPI) Delete(c *gin.Context) {
	var body models.GetPassword

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

	_, err := passService.Delete(body.Name, id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, "Удаление прошло успешно")
}
