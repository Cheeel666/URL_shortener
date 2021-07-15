package server

import (
	"URL_shortener/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	//"URL_shortener/base62"
)

// Создаем сервер и эндпоинты
func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})
	r.POST("/short", utils.Shortener)
	r.POST("/long", utils.Longener)

	return r
}
