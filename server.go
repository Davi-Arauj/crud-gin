package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/teste", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK!!",
		})
	})
	server.Run(":8080")

}
