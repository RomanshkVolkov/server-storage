package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/RomanshkVolkov/server-storage/internal/adapters/repository"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now().UTC()
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		user, err := repository.ExtractDataByToken(token)

		if err == nil && user.ID != 0 {
			c.Set("user", user)
		}

		origin := "*" //c.Request.Header.Get("Origin") // Get the origin of the request
		fmt.Println(origin)
		if origin != "" { // Only set CORS headers if there's an Origin header
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin) // Reflect the origin

			if c.Request.Method == "OPTIONS" {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, ngrok-skip-browser-warning")
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Important for credentials!
				c.Writer.Header().Set("Access-Control-Max-Age", "86400")          // Cache preflight response for 24 hours

				c.AbortWithStatus(http.StatusNoContent) // 204 No Content is the standard response for OPTIONS
				return
			}
		} else {
			// Handle requests without Origin header (e.g., server-side) - likely not needed for your case
			// You might want to log this or handle it differently.
			fmt.Println("Request without Origin header")
		}

		c.Next()

		latency := time.Since(t)
		fmt.Println(latency)

		status := c.Writer.Status()
		fmt.Println(status)

	}
}
