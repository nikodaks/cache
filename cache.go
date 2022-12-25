package cache

import (
	"sync"
	"time"
)

type Cache struct {
	values          map[string]*cacheItem
	mu              *sync.RWMutex
	cleanUpInterval time.Duration
}

func (c *Cache) CountCacheItems() int {
	return len(c.values)
}

func (c *Cache) Set(key string, value interface{}, timeToDelete time.Duration) {
	cacheItem := newCacheItem(value, timeToDelete)
	c.values[key] = cacheItem
}

func (c *Cache) Get(key string) (interface{}, error) {
	_, exists := c.values[key]

	if !exists {
		return nil, c.notFoundItemError(key)
	}

	return c.values[key], nil
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.values[key]

	if !exists {
		return c.notFoundItemError(key)
	}

	delete(c.values, key)

	return nil
}

func (c *Cache) CleanUp() {
	for {
		c.cleanUpValues()
		<-time.After(c.cleanUpInterval)
	}
}

func (c *Cache) cleanUpValues() {
	if len(c.values) == 0 {
		return
	}

	jobs := make(chan string)

	go c.cleanUpWorker(jobs)

	for key := range c.values {
		jobs <- key
	}

	close(jobs)
}

func (c *Cache) cleanUpWorker(jobs <-chan string) {
	for cacheItemKey := range jobs {
		cacheItem := c.values[cacheItemKey]
		timeNow := time.Now().Unix()
		shouldClearData := timeNow >= cacheItem.timeToDelete

		if shouldClearData {
			c.Delete(cacheItemKey)
		}
	}
}

func New(shouldClearData bool, cleanUpInterval time.Duration) *Cache {
	c := &Cache{
		values:          make(map[string]*cacheItem),
		mu:              &sync.RWMutex{},
		cleanUpInterval: cleanUpInterval,
	}

	if shouldClearData {
		go c.CleanUp()
	}

	return c
}
