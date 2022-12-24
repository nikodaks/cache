<h1>In memory cache</h1>

<h2>Example of usage:</h2>

```go
package main

import (
  "fmt"
  "github.com/nikodaks/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}

```