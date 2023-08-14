package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type User struct {
	Id       int64  `redis:"id"`
	Name     string `redis:"name"`
	Email    string `redis:"email"`
	Password string `redis:"password"`
}

func main() {
	ctx := context.Background()
	user := &User{
		Id:       1,
		Name:     "Jake",
		Email:    "arneg0shua@gmail.com",
		Password: "password",
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6384",
		Password: "",
		DB:       0,
	})
	if err := rdb.HSet(ctx, "users:1", user).Err(); err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
	output, err := rdb.HGetAll(ctx, "users:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
