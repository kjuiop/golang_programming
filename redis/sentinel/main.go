package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// Redis Sentinel 주소 설정
	sentinelAddrs := []string{fmt.Sprintf("%s:26379", "127.0.0.1")}

	// Redis 클라이언트 생성
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "kollus",         // Sentinel에서 사용하는 마스터 이름
		SentinelAddrs: sentinelAddrs,    // Sentinel 노드 주소
		DB:            0,                // 사용할 데이터베이스 번호
		Password:      "",               // Redis 비밀번호 (없는 경우 빈 문자열)
		PoolSize:      10,               // 풀 크기
		MinIdleConns:  3,                // 최소 유지 커넥션 수
		ReadTimeout:   5 * time.Second,  // 읽기 타임아웃
		WriteTimeout:  5 * time.Second,  // 쓰기 타임아웃
		IdleTimeout:   30 * time.Second, // 유휴 타임아웃
	})

	// Redis 연결 확인
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis 연결 확인:", pong)

	// 데이터 쓰기
	err = client.Set(context.Background(), "example_key", "Hello, Redis Sentinel!", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// 데이터 읽기
	val, err := client.Get(context.Background(), "example_key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("example_key: %s\n", val)

	// Redis 클라이언트 종료
	if err := client.Close(); err != nil {
		log.Fatal(err)
	}
}

func getContainerIP(networkName, containerName string) (string, error) {
	// Docker 네트워크에서 컨테이너의 IP 주소를 얻기
	cmd := fmt.Sprintf("docker network inspect -f '{{range .Containers}}{{if eq .Name \"%s\"}}{{.IPv4Address}}{{end}}{{end}}' %s", containerName, networkName)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}

	// 얻은 IP 주소에서 '/24' 부분 제거
	ip := strings.Split(string(output), "/")[0]

	return ip, nil
}
