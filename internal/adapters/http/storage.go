package http

import (
	"github.com/RomanshkVolkov/server-storage/internal/adapters/handler"
	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.Engine) {
	storage := r.Group("/storage")
	{
		storage.POST("", handler.UploadFile)
		storage.DELETE("/:path", handler.DeleteFile)
	}
}
