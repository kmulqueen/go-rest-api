package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmulqueen/go-rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized."})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized."})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
