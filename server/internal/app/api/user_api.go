package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/app/services"
	"server/internal/models"
	"server/internal/tools"
)

// UserAPI структура
type UserAPI struct{}

// NewUserAPI метод возвращает указатель на структуру UserAPI со всеми ее методами
func NewUserAPI() *UserAPI {
	return &UserAPI{}
}

var userService = services.NewUserService()

// Save сохраняет пользователя
func (ua *UserAPI) Save(c *gin.Context) {
	var body models.SaveUserRequest

	if err := tools.RequestBinderBody(&body, c); err != nil {
		return
	}

	userModel, err := userService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, userModel)
}

// Get возвращает пользователя
func (ua *UserAPI) Get(c *gin.Context) {
	userModel, err := userService.Get()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, userModel)
}
