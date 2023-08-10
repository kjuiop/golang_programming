package main

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
	"time"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	cache.Set("key", "value", 5)
	cache.SetWithTTL("key2", "Good Morning", 12, 10*time.Minute)

	time.Sleep(10 * time.Millisecond)

	value, found := cache.Get("key")
	if !found {
		panic("missing value")
	}
	fmt.Println(value)
	cache.Del("key")
}
