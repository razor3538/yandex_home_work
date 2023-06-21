package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/internal/app/api"
	middleware "server/internal/app/midleware"
)

// SetupRouter setting up gin router and config
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20

	r.Static("/static", "./static")

	user := api.NewUserAPI()
	pass := api.NewPasswordAPI()
	text := api.NewTextAPI()
	card := api.NewCardAPI()
	auth := api.NewAuth()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/api/user/register", user.Save)

	r.POST("/api/user/login", middleware.Passport().LoginHandler)
	r.GET("/api/user/is-authenticated", auth.IsAuthenticated)

	authRequired := r.Group("/")
	authRequired.Use(middleware.JwtAuthMiddleware())
	{
		r.POST("/api/passwords", pass.Save)
		r.POST("/api/get_password", pass.Get)
		r.DELETE("/api/password", pass.Delete)

		r.POST("/api/text", text.Save)
		r.POST("/api/get_text", text.Get)
		r.DELETE("/api/text", text.Delete)

		r.POST("/api/card", card.Save)
		r.POST("/api/get_card", card.Get)
		r.DELETE("/api/card", card.Delete)
	}

	return r
}
