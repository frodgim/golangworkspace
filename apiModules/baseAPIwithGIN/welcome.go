package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterWelcomeAPI(router *gin.Engine) {
	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API!",
		})
	})
}
