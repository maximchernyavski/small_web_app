package middlewares

import (
	"net/http"

	"example.com/web_shit/utils"
	"github.com/gin-gonic/gin"
)

func Verify(context *gin.Context) {
	token := context.Request.Header.Get("token")

	_, isAdmin, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't verify token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User is verified", "isAdmin": isAdmin})
}
