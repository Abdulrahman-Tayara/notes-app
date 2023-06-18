package api

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/http"
	controllers2 "github.com/Abdulrahman-Tayara/notes-app/users-service/api/controllers"
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

	apiGroup.POST("/signup", http.GinWrapper(controllers2.SignUpController))
	apiGroup.POST("/login", http.GinWrapper(controllers2.LoginController))

	apiGroup.POST("/refresh", http.GinWrapper(controllers2.RefreshAccessToken))
	apiGroup.POST("/logout", http.GinWrapper(controllers2.Logout))
}