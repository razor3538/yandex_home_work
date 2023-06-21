package middleware

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	repositories "server/internal/app/repository"
	"server/internal/models"
	"server/internal/tools"
	"time"
)

// JwtAuthMiddleware мидлвейр авторизации
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tools.TokenValid(c)
		if err != nil {
			println(err.Error())
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

// IdentityKeyID is used to tell
// by what field we will identify user
const IdentityKeyID = "id"

// UserID struct
type UserID struct {
	ID string
}

var userRepo = repositories.NewUserRepo()

// Passport is middleware for user authentication
func Passport() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "YandexPracticum",
		Key:         []byte(os.Getenv("JWTSECRET")),
		Timeout:     time.Hour * 4,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: IdentityKeyID,
		TokenLookup: "header:Authorization",
		LoginResponse: func(c *gin.Context, i int, s string, t time.Time) {
			value, _ := Passport().ParseTokenString(s)
			id := jwt.ExtractClaimsFromToken(value)["id"]
			result, err := userRepo.GetByKey("id", id.(string))

			if err != nil {
				tools.CreateError(http.StatusUnauthorized, err, c)
				return
			}

			c.Header("token", value.Raw)
			c.JSON(http.StatusOK, result)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserID); ok {
				return jwt.MapClaims{
					IdentityKeyID: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserID{
				ID: claims[IdentityKeyID].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var body models.LoginUserModel
			if err := c.ShouldBind(&body); err != nil {
				return "", errors.New("пароль или логин не введен")
			}

			result, err := userRepo.GetByKey("login", body.Username)

			if err != nil {
				return nil, errors.New("не верный логин или пароль")
			}
			equal := tools.CheckPasswordHash(body.Password, result.Password)
			if equal {
				return &UserID{
					ID: result.ID.String(),
				}, nil
			}
			return nil, errors.New("не верный логин или пароль")
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})
	return authMiddleware
}
