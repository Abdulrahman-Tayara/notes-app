package http

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/api/http/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	nethttp "net/http"
)

func SetupRouters(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	engine.Use(cors.New(corsConfig))

	apiGroup := engine.Group("api/")

	apiGroup.GET("/health", func(context *gin.Context) {
		context.JSON(nethttp.StatusOK, gin.H{
			"message": "I'm good thanks for asking",
		})
	})

	apiGroup.POST("/signup", http.GinWrapper(controllers.SignUpController))
	apiGroup.POST("/login", http.GinWrapper(controllers.LoginController))

	apiGroup.POST("/refresh", http.GinWrapper(controllers.RefreshAccessToken))
	apiGroup.POST("/logout", http.GinWrapper(controllers.Logout))
}
