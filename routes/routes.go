package routes

import (
	"example.com/web_shit/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	server.POST("/auth", auth)
	server.POST("/verify", middlewares.Verify)
	server.POST("/signup", signup)
}
