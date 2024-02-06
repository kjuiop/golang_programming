package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	// Redis 클라이언트 생성
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 서버 주소
		Password: "",               // Redis 비밀번호 (없는 경우 빈 문자열)
		DB:       0,                // 사용할 데이터베이스 번호
	})

	// Redis 연결 확인
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis 연결 확인:", pong)

	// 데이터 쓰기
	err = client.Set(context.Background(), "example_key", "Hello, Redis!", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// 데이터 읽기
	val, err := client.Get(context.Background(), "example_key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("example_key: %s\n", val)

	// 특정 키의 만료 시간 설정
	err = client.Expire(context.Background(), "example_key", 10*time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}

	// 잠시 대기 후 데이터 확인 (만료 확인)
	time.Sleep(2 * time.Second)
	val, err = client.Get(context.Background(), "example_key").Result()
	if err == redis.Nil {
		fmt.Println("example_key: 만료됨")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("example_key: %s\n", val)
	}

	// Redis 클라이언트 종료
	if err := client.Close(); err != nil {
		log.Fatal(err)
	}
}
