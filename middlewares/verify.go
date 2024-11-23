package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/web_shit/utils"
	"github.com/gin-gonic/gin"
)

func Verify(context *gin.Context) {
	token := context.Request.Header.Get("token")
	isAdmin, err := strconv.ParseBool(context.Request.Header.Get("isAdmin"))
	fmt.Println("isAdmin before parse ===>", context.Request.Header.Get("isAdmin"))
	fmt.Println("First isAdmin ===>", isAdmin)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse admin state"})
		return
	}
	fmt.Println("TokenBack ===>", token)
	fmt.Println("isAdmin back ===>", isAdmin)

	_, retrievedIsAdmin, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't verify token"})
		return
	}

	if isAdmin != retrievedIsAdmin {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Wrong admin state"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User is verified", "isAdmin": isAdmin})
}
