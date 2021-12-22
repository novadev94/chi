package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	deployTime := time.Now()
	router := gin.Default()
	router.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"deployTime": deployTime,
			"timestamp":  time.Now(),
		})
	})
	return &Server{
		router: router,
	}
}

func (a *Server) Serve(port string) error {
	return a.router.Run(fmt.Sprintf(":%s", port))
}
