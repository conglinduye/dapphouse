package middleware

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"dapphouse/handler"
	"dapphouse/common/token"
	"dapphouse/common/errno"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			fmt.Printf("AuthMiddleware error: %s\n", err)
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
