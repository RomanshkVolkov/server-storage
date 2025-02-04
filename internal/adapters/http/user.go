package http

import (
	"github.com/RomanshkVolkov/server-storage/internal/adapters/handler"
	"github.com/RomanshkVolkov/server-storage/internal/adapters/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	protect := middleware.Protected
	users := r.Group("/users")
	{
		users.GET("/", protect(), handler.GetAllUsers)
		users.POST("/", protect(), handler.CreateUser)
		users.PUT("/:id", protect(), handler.UpdateUser)
		users.DELETE("/:id", protect(), handler.DeleteUser)

	}

}
