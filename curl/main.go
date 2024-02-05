package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 대상 URL
	url := "https://localhost:8080"

	// 종료 신호 처리
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// 무한 루프
	for {
		fmt.Printf("Request (infinite loop):\n")
		makeGetRequest(url)
		fmt.Println("--------------------------------------")

		// 종료 신호를 받으면 루프 종료
		select {
		case <-stopChan:
			fmt.Println("Received termination signal. Exiting.")
			return
		default:
			// 계속 진행
		}

		// 적절한 딜레이를 넣어도 좋습니다.
		time.Sleep(3 * time.Second) // 1초 딜레이 (원하는 시간으로 변경 가능)
	}
}

func makeGetRequest(url string) {
	// GET 요청 보내기
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
	}

	// 응답 출력
	fmt.Printf("Res Data : %s\n", string(body))
	// 여기에서 필요한 응답 처리를 추가할 수 있습니다.
}
