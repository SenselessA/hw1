# Реализация TTL

```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SenselessA/hw1/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*5)

	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Println(err)
	}
	fmt.Println(userId) // Output: 42

	time.Sleep(time.Second * 6) // прошло 5 секунд

	_, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}
}
```