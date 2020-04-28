package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello there, this is a test application for Kubernetes! "+
			"This application has been changed in the GitHub master branch, built in Docker Hub and then updated automatically by Kubernetes!")
	})

	r.Run(":8080")
}
