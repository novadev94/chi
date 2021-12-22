package server

import (
	"fmt"
	"github.com/haitien/chi/api"
	"github.com/haitien/chi/inject"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	deployTime := time.Now()
	engine := gin.Default()
	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"deployTime": deployTime,
			"timestamp":  time.Now(),
		})
	})
	server := &Server{
		engine: engine,
	}
	server.registerRoutes()
	return server
}

func (a *Server) registerRoutes() {
	injector := inject.NewInjector()
	// TODO provide services then register routes
	service := injector.ProvideService()
	controller := api.NewApi(service)
	v1 := a.engine.Group("v1")
	v1.GET("tokenPrice", controller.TokenPrice)
}

func (a *Server) Serve(port string) error {
	return a.engine.Run(fmt.Sprintf(":%s", port))
}
