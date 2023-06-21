package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	middleware "server/internal/app/midleware"
	repositories "server/internal/app/repository"
	"server/internal/domain"
)

// AuthService структура
type AuthService struct{}

// NewAuthService метод возвращает указатель на структуру AuthService со всеми ее методами
func NewAuthService() *AuthService {
	return &AuthService{}
}

var userRepo = repositories.NewUserRepo()

// IsAuthenticated метод проверяет авторизирован ли пользователь
func (as *AuthService) IsAuthenticated(c *gin.Context) (domain.User, int, error) {
	claims, err := middleware.Passport().CheckIfTokenExpire(c)

	if err != nil {
		return domain.User{}, http.StatusUnauthorized, err
	}
	if int64(claims["exp"].(float64)) < middleware.Passport().TimeFunc().Unix() {
		_, _, _ = middleware.Passport().RefreshToken(c)
	}

	id := claims[middleware.IdentityKeyID]
	result, err := userRepo.GetByKey("id", id.(string))

	if err != nil {
		return domain.User{}, http.StatusUnauthorized, errors.New("токен не действителен")
	}
	return result, 0, nil
}
