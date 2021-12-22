package api

import (
	"github.com/gin-gonic/gin"
	"github.com/haitien/chi/service"
	"net/http"
)

type MyApi struct {
	service *service.Service
}

func NewApi(service *service.Service) *MyApi {
	return &MyApi{
		service: service,
	}
}

type TokenPriceOutput struct {
	Message string `json:"message"`
}

func (a *MyApi) TokenPrice(c *gin.Context) {
	err := a.service.GenData("1m")
	if err != nil {
		c.JSON(http.StatusForbidden, &TokenPriceOutput{
			Message: "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, &TokenPriceOutput{
		Message: "Success",
	})
}
