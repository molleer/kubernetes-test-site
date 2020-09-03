package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is a test site running on nginx. The the application rebuilds when master is updated on github. This is really cool and its really fast!\n"+
			"This repo has also been used to debug molleer/blame-alert!")
	})

	r.Run(":8080")
}
