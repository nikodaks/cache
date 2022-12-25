<h1>In memory cache</h1>

<h2>Example of usage:</h2>

```go
package main

func main() {
	key := "1"
	value := 2
	timeToLive := time.Second * 2

	CleanupInterval := time.Second * 1
	startCleanup := true 

	c := GoCache.NewCache(CleanupInterval, startCleanup)
	c.Set(key, value, timeToLive)

	_, err := c.Get(key)
	if err == nil {
		fmt.Println("found element")

	}

	// wait for cleanup data
	time.Sleep(time.Second * 5)

	v, err := c.Get("1")
	if err == nil {
		fmt.Printf("found element %v \n", v)
	} else {
		fmt.Println(err)
	}

	ok := c.Delete(key)
	if ok {
		fmt.Printf("Delete %s, number of elements in the cache %d", key, c.CountOfElement())
	} else {
		fmt.Printf("not found %s, number of elements in the cache %d", key, c.CountOfElement())
	}
}

```