package utils

import (
	"URL_shortener/database"
	"URL_shortener/app/base62"
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
	json.Unmarshal([]byte(jsonData), &url)

	result, err := NewShortUrl(url.Url)
	if err != nil  || len(url.Url) == 0{
		c.AbortWithStatus(500)
	} else {
		c.JSON(200, gin.H{"url":result})
	}
}

func Longener(c *gin.Context)  {
	var url UrlStruct

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal([]byte(jsonData), &url)

	result, err := GetLongUrl(url.Url)

	if err != nil  || len(url.Url) == 0 || len(result) == 0{
		c.AbortWithStatus(500)
	} else {
		c.JSON(200, gin.H{"url":result})
	}
}

func NewShortUrl(url string) (string, error) {
	id, err := database.AddUrl(url)
	if err != nil {
		return "", err
	}
	shorterUrl := base62.ToAlphabet(id - 1)
	database.UpdateUrl(id, shorterUrl)
	return shorterUrl, nil
}

func GetLongUrl(url string) (string, error) {
	result, err := database.SelectLongUrl(url)
	if err != nil {
		return "", err
	}
	return result, nil
}