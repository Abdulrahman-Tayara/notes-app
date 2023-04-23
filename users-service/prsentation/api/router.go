package api

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/controllers"
	"github.com/gin-gonic/gin"
	nethttp "net/http"
)

func SetupRouters(engine *gin.Engine) {
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
