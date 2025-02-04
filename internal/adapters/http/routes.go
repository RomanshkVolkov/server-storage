package http

import (
	"net/http"

	"github.com/RomanshkVolkov/server-storage/internal/adapters/handler"
	"github.com/RomanshkVolkov/server-storage/internal/adapters/middleware"
	"github.com/RomanshkVolkov/server-storage/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// 404 route

	r.NoRoute(notFound)

	r.Use(middleware.Middleware())

	r.Static("/files", "/srv/files")

	AuthRoutes(r)
	UserRoutes(r)
	MailRoutes(r)
	StorageRoutes(r)

	// root route
	r.GET("/", handler.HomeResponse)

}

func notFound(c *gin.Context) {
	notFountResponse := domain.APIResponse[any]{
		Success: false,
		Message: domain.Message{
			En: "Route not found",
			Es: "Ruta no encontrada",
		},
	}

	c.IndentedJSON(http.StatusNotFound, notFountResponse)
}
