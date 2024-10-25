package routes

import (
	"github.com/gin-gonic/gin"
)

func getMain(context *gin.Context) {
	context.File("/pages/main.html")
}
