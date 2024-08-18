package network

import "github.com/gin-gonic/gin"

func response(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
	return
}
