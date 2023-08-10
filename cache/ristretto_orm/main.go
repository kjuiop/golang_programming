package main

import (
	"fmt"
	"golang_programming/cache/ristretto_orm/models"
	"time"
)

func main() {
	engn, err := models.NewEngn()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := engn.CloseEngn()
		if err != nil {
			panic(err)
		}
	}()

	user1, fromCache1, err := engn.FindUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fromCache:%t, user:%#v\n", fromCache1, user1)
	time.Sleep(100 * time.Millisecond)

	user2, fromCache2, err := engn.FindUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fromCache:%t, user:%#v\n", fromCache2, user2)
}
