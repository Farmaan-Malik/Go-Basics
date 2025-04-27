package middleware

import (
	"example/rest-api/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"status": "failure", "error": "missing auth header"})
		return
	}
	id, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	ctx.Set("userId", id)
	ctx.Next()
}
