package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Gin 엔진을 초기화하고 기본 설정을 사용합니다.
	gMux := gin.Default()

	// HTML 템플릿 로드
	gMux.LoadHTMLGlob("./templates/*")

	// 루트 경로에 대한 핸들러
	gMux.GET("/", func(c *gin.Context) {
		// index.html 템플릿을 렌더링하여 응답합니다.
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to My Website!",
		})
	})

	// 서버 시작
	_ = gMux.Run(":8080")
}
