package middlewares

import (
	"net/http"

	"github.com/LKezHn/React-Go-Project/services"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token no provided",
			})
			c.Abort()
			return
		}

		id, err := services.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", id)
		c.Next()
	}
}
