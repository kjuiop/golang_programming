package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	UsePolicy bool   `json:"use_policy"`
}

func main() {
	// Gin 라우터 생성
	router := gin.Default()

	// 라우터 정의
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "안녕하세요!",
		})
	})

	// POST 요청을 처리하는 핸들러
	router.POST("/user", func(c *gin.Context) {
		// JSON 요청 바인딩
		var user User

		user.UsePolicy = true
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println("policy : ", user.UsePolicy)

		// 받은 데이터 출력
		c.JSON(http.StatusOK, gin.H{
			"message":                   "유저 정보를 받았습니다.",
			"name":                      user.Name,
			"email":                     user.Email,
			"UseNicknameDuplicateBlock": user.UsePolicy,
		})
	})

	// Gin 서버 시작
	router.Run(":3000")
}
