package utils

import "github.com/gin-gonic/gin"
import "URL_shortener/base62"
import "URL_shortener/db"

func shortener(c *gin.Context)  {
	var result_url string
	result_url = "Ok"
	c.JSON(200, gin.H{"url1":result_url})
}

func longener(c *gin.Context)  {
	var result_url string
	result_url = "Ok"
	c.JSON(200, gin.H{"url1":result_url})
}

func NewShortUrl(url string) string {
	id := db.AddUrl(url)
	shorterUrl := base62.ToAlphabet(id)
	db.UpdateUrl(id, shorterUrl)
	return shorterUrl
}
