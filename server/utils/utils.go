package utils

import (
	"URL_shortener/base62"
	"URL_shortener/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)


type UrlStruct struct {
	Url string
}


// хендлеры для пост запрсоов
func Shortener(c *gin.Context)  {
	var url UrlStruct

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(jsonData), &url)

	c.JSON(200, gin.H{"url":NewShortUrl(url.Url)})

}

func Longener(c *gin.Context)  {
	var url UrlStruct

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(jsonData), &url)

	c.JSON(200, gin.H{"url":GetLongUrl(url.Url)})
}

func NewShortUrl(url string) string {
	id := db.AddUrl(url)
	shorterUrl := base62.ToAlphabet(id - 1)
	db.UpdateUrl(id, shorterUrl)
	return shorterUrl
}

func GetLongUrl(url string) string {
	return db.SelectLongUrl(url)
}