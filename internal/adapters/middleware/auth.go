package middleware

import (
	"net/http"

	"github.com/RomanshkVolkov/server-storage/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")

		if !exists || user == nil {
			c.IndentedJSON(http.StatusUnauthorized, domain.APIResponse[domain.RequestInfo]{
				Success: false,
				Message: domain.Message{
					En: "Unauthorized",
					Es: "No autorizado",
				},
				Data: domain.RequestInfo{
					Host:      c.Request.Host,
					IP:        c.Request.RemoteAddr,
					UserAgent: c.Request.UserAgent(),
					UserID:    0,
				},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
