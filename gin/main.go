package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	r.Run(":3020")
}
