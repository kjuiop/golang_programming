package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	graceful "gopkg.in/tylerb/graceful.v1"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

	gMux.LoadHTMLGlob("./templates/*")
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

			if len(part.FileName()) == 0 {
				break
			}

			filename := filepath.Join("./result/", part.FileName())
			out, err := os.Create(filename)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Create file error: %s", err.Error()))
				return
			}
			defer out.Close()

			var (
				read    int64
				written int64
				next    int32 = 5
			)

			length := c.Request.ContentLength
			buffer := make([]byte, 1024*1024)
			for {
				bytes, readErr := part.Read(buffer)
				if readErr == io.EOF {
					break
				}

				if readErr != nil {
					c.String(http.StatusInternalServerError, fmt.Sprintf("file byte read error: %s", err.Error()))
					return
				}

				if bytes <= 0 || buffer == nil {
					c.String(http.StatusInternalServerError, fmt.Sprintf("close upload session: %s", err.Error()))
					return
				}

				read = read + int64(bytes)
				writtenBytes, err := out.Write(buffer[0:bytes])
				if err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprintf("write session: %s", err.Error()))
					return
				}
				written = written + int64(writtenBytes)

				p := int32(float32(read) / float32(length) * 100)
				if 0 != p && 0 == p%next {
					log.Printf("upload processing : %s\n", strconv.Itoa(int(p)))
					next += 5
				}

				if p == 99 {
					log.Printf("upload processing complete")
					break
				}

			}

			_, err = io.Copy(out, part)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Copy file error: %s", err.Error()))
				return
			}

			c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", part.FileName()))
		}

	})

	if *uploadMode == "general" {
		gMux.Run(":8082")
	} else if *uploadMode == "graceful" {
		srv := &graceful.Server{
			Timeout:   0,
			ConnState: func(conn net.Conn, state http.ConnState) {},
			Server: &http.Server{
				Addr:    ":8082",
				Handler: gMux,
			},
		}
		_ = srv.ListenAndServe()
	}

}
