package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/RomanshkVolkov/server-storage/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func HomeResponse(c *gin.Context) {
	req := c.Request
	var userID uint
	id, exist := c.Get("userID")

	if exist {
		userID = id.(uint)
	}

	c.JSON(http.StatusOK, domain.APIResponse[domain.RequestInfo]{
		Success: true,
		Message: domain.Message{
			En: "Welcome to the API",
			Es: "Bienvenido a la API",
		},
		Data: domain.RequestInfo{
			Host:      req.Host,
			IP:        req.RemoteAddr,
			UserAgent: req.UserAgent(),
			UserID:    userID,
		},
	})

}

var RequestError = domain.Message{
	En: "Request error",
	Es: "Error en la solicitud",
}

func ValidateRequest[T any](c *gin.Context) (*T, error) {
	var request T
	if err := c.BindJSON(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func ServerError(err error, message domain.Message) domain.APIResponse[any] {
	response := domain.APIResponse[any]{
		Success: false,
		Message: message,
		Data:    nil,
		Error:   err,
	}
	return response
}

func ExtractAndParseUintParam(c *gin.Context, paramName string) (uint, error) {
	stringParam := c.Param(paramName)
	id, err := strconv.ParseUint(stringParam, 10, 64)
	if err != nil {
		InvalidParamError(c, paramName, err)
	}

	return uint(id), err
}

func ExtractQueryParam(c *gin.Context, paramName string) string {
	return c.Query(paramName)
}

func ExtractAndParseUintQueryParam(c *gin.Context, paramName string) (uint, error) {
	stringParam := ExtractQueryParam(c, paramName)
	id, err := strconv.ParseUint(stringParam, 10, 64)
	if err != nil {
		InvalidParamError(c, paramName, err)
	}

	return uint(id), err
}

func InvalidParamError(c *gin.Context, paramName string, err error) {
	c.IndentedJSON(http.StatusBadRequest, ServerError(err, domain.Message{
		En: "Invalid parameter " + paramName,
		Es: "Parámetro inválido " + paramName,
	}))
}

func GetSubdomain(c *gin.Context) string {
	host := c.Request.Host
	splitedHost := strings.Split(host, ".")

	if len(splitedHost) > 2 {
		return splitedHost[0]
	}

	return ""
}
