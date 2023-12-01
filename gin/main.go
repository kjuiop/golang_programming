package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var uploadMode *string

func main() {

	uploadMode = flag.String("mode", "", "upload mode, graceful or general -mode {mode}")
	flag.Parse()

	if *uploadMode == "" {
		flag.Usage() // 사용법 출력
		log.Fatalf("fail init, require upload mode, graceful or general -mode {mode} flag value")
	}

	gMux := gin.Default()

	gMux.LoadHTMLGlob("/home/jake/workspace/study/golang_programming/gin/templates/*")
	gMux.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	gMux.POST("/upload", func(c *gin.Context) {
		// MultipartReader로부터 form 데이터를 읽기
		multipartReader, err := c.Request.MultipartReader()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("MultipartReader error: %s", err.Error()))
			return
		}

		for {
			part, err := multipartReader.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("NextPart error: %s", err.Error()))
				return
			}

			// 업로드된 파일 처리
			if part.FileName() != "" {
				filename := filepath.Join("/home/jake/workspace/study/golang_programming/gin/uploads", part.FileName())
				out, err := os.Create(filename)
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprintf("Create file error: %s", err.Error()))
					return
				}
				defer out.Close()

				_, err = io.Copy(out, part)
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprintf("Copy file error: %s", err.Error()))
					return
				}

				c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", part.FileName()))
			}
		}

	})

	if *uploadMode == "general" {
		gMux.Run(":8081")
	} else if *uploadMode == "graceful" {
		srv := &graceful.Server{
			Timeout:   0,
			ConnState: func(conn net.Conn, state http.ConnState) {},
			Server: &http.Server{
				Addr:    ":8081",
				Handler: gMux,
			},
		}

		_ = srv.ListenAndServe()
	}

}
