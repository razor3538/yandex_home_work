package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/app/services"
	"server/internal/tools"
)

// Auth структура
type Auth struct{}

// NewAuth метод возвращает указатель на структуру Auth со всеми ее методами
func NewAuth() *Auth {
	return &Auth{}
}

var authService = services.NewAuthService()

// IsAuthenticated метод проверяющий авторизирован ли пользователь
func (a Auth) IsAuthenticated(c *gin.Context) {
	user, code, err := authService.IsAuthenticated(c)

	if err != nil {
		tools.CreateError(code, err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

// login
func login() {}

// logout
func logout() {}
