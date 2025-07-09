package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	auth = "Bearer "
)

func (i *Implementation) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token required",
			})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, auth)
		access, err := i.authService.Check(c.Request.Context(), token)
		if err != nil || !access {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token invalid",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
