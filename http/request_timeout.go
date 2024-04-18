package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 대상 URL
	url := "http://localhost:3000/context"
	log.Println("----------------------------------start")
	makeGetRequest(url)
	log.Println("----------------------------------terminal")
}

func makeGetRequest(url string) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// HTTP 클라이언트 생성
	client := http.DefaultClient

	// HTTP 요청 실행
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 응답 출력
	fmt.Println("Response Status:", resp.Status)
}
