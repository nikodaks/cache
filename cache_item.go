package cache

import "time"

type cacheItem struct {
	value        interface{}
	timeToDelete int64
}

func newCacheItem(value interface{}, liveDuration time.Duration) *cacheItem {
	return &cacheItem{value: value, timeToDelete: time.Now().Add(liveDuration).Unix()}
}
