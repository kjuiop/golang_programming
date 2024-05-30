package main

import (
	"log"
	"os"
	"time"
)

func main() {
	tick := time.Tick(time.Duration(5) * time.Second)

	for {
		select {
		case <-tick:

			dir := "/home/jake/studyspace/go/golang_programming/time/tick/movie"
			list, err := os.ReadDir(dir)
			if err != nil {
				log.Printf("read dir, err : %v\n", err)
				return
			}
			log.Printf("dir info : %v\n", list)
			for _, tempDir := range list {
				if !tempDir.IsDir() {
					continue
				}

				fileInfo, err := tempDir.Info()
				if err != nil {
					log.Println("Error:", err)
					continue
				}

				log.Printf("fileInfo : %v\n", fileInfo)

				// 파일 또는 디렉토리의 수정 시간 가져오기
				modTime := fileInfo.ModTime()

				log.Printf("file info mod time : %v\n", modTime)
			}

			log.Println("hello tick tock")
		default:
			time.Sleep(time.Second)
		}
	}
}
