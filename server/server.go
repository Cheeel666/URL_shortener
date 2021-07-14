package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"URL_shortener/server/utils"
	//"URL_shortener/base62"
)

type UrlStruct struct {
	Url string
}


// хендлеры для пост запрсоов
func shortener(c *gin.Context)  {
	var result_url string
	var url Url

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic("Error during parsing JSON")
	}

	json.Unmarshal([]byte(jsonData), &url)

	result_url = utils.NewShortUrl(url.Url)
	c.JSON(200, gin.H{"url1":result_url})

}

func longener(c *gin.Context)  {
	var result_url string
	result_url = "Ok"
	c.JSON(200, gin.H{"url1":result_url})
}

// Создаем сервер и эндпоинты
func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})
	r.POST("/short", shortener)
	r.POST("/long", longener)

	return r
}
