package main

import (
	"cache/cache"
	"fmt"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId2 := cache.Get("userId")

	fmt.Println(userId2)
}
