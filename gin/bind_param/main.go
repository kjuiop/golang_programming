package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	HardwareID string `form:"hardware_id"`
	SessionKey string `form:"session_key"`
	PlayerID   string `form:"player_id"`
	UserID     string `form:"user_id"`
	ContentKey string `form:"content_key"`
	BlockKey   string `form:"block_key"`
}

func main() {
	router := gin.Default()

	router.GET("/check", func(c *gin.Context) {
		var params RequestParams
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("json : %v", params)

		c.JSON(200, gin.H{"status": "success"})
	})

	router.Run(":8080")
}
