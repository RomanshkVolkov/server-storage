package handler

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/RomanshkVolkov/server-storage/internal/core/domain"
	"github.com/RomanshkVolkov/server-storage/internal/core/service"
	"github.com/gin-gonic/gin"
)

type SomeRequest struct {
	FileList []multipart.FileHeader `form:"files" binding:"required"`
}

func UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ServerError(err, RequestError))
		return
	}

	files := form.File["files"]

	if len(files) == 0 {
		c.IndentedJSON(http.StatusBadRequest, ServerError(nil, domain.Message{En: "no files uploaded", Es: "no se subieron archivos"}))
		return
	}

	baseDir := "/srv/files"
	host := c.Request.Host
	createdFiles := []domain.StorageCreatedFile{}
	invalidFiles := []domain.InvalidFile{}

	for _, file := range files {
		// 2MB equal to 2 * 1024
		invalidSize, message := service.ValidateSizeFile(file, 2*1024*1024)
		if !invalidSize {
			invalidFiles = append(invalidFiles, message)
			continue
		}
		isImage := service.ValidateFileIsImage(file)
		subDir := "/documents"
		if isImage {
			subDir = "/images"
		}

		dirPath := filepath.Join(baseDir, subDir)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, ServerError(err, domain.Message{En: "Error creating folder", Es: "Error creando carpeta"}))
			invalidFiles = append(invalidFiles, domain.InvalidFile{
				Name:    file.Filename,
				Message: "Error creating folder",
			})
			continue
		}

		filePath := filepath.Join(dirPath, file.Filename)
		fileAbsoluteOutputPath := strings.TrimSuffix(filePath, filepath.Ext(filePath))
		fileOutputPath := strings.Replace(fileAbsoluteOutputPath, "/srv", "", 1)

		if isImage {
			webpPath := fileAbsoluteOutputPath + ".webp"

			if err := service.ConvertToWebP(file, webpPath); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, ServerError(err, domain.Message{En: "error converting to webp", Es: "error convirtiendo a webp"}))
				invalidFiles = append(invalidFiles, domain.InvalidFile{
					Name:    file.Filename,
					Message: "Error converting to webp",
				})
				continue
			}

			createdFiles = append(createdFiles, domain.StorageCreatedFile{
				FullPath: "https://" + host + fileOutputPath + ".webp",
				Path:     fileOutputPath + ".webp",
				Folder:   subDir,
				Name:     strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)) + ".webp",
				Mime:     "image/webp",
			})
		} else {
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, ServerError(err, domain.Message{En: "error saving file", Es: "error guardando archivo"}))
				invalidFiles = append(invalidFiles, domain.InvalidFile{
					Name:    file.Filename,
					Message: "Error saving file",
				})
			}

			filePath = fileOutputPath + filepath.Ext(filePath)
			createdFiles = append(createdFiles, domain.StorageCreatedFile{
				FullPath: "https://" + host + filePath,
				Path:     filePath,
				Folder:   subDir,
				Name:     file.Filename,
				Mime:     file.Header.Get("Content-Type"),
			})
		}
	}
	c.IndentedJSON(http.StatusOK, domain.APIResponse[domain.UploadedResponse]{
		Success: true,
		Message: domain.Message{
			En: "files uploaded",
			Es: "archivos subidos",
		},
		Data: domain.UploadedResponse{
			Upload:       createdFiles,
			InvalidFiles: invalidFiles,
		},
	})
}

func DeleteFile(c *gin.Context) {
	filePath := c.Param("path")
	fullPath := filepath.Join("/srv", filePath)

	if err := os.Remove(fullPath); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ServerError(err, domain.Message{En: "error deleting file", Es: "error eliminando archivo"}))
		return
	}

	c.IndentedJSON(http.StatusOK, domain.APIResponse[any]{
		Success: true,
		Message: domain.Message{
			En: "file deleted",
			Es: "archivo eliminado",
		},
		Data: nil,
	})
}
