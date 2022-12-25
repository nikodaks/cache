package cache

import "fmt"

func (c *Cache) notFoundItemError(key string) error {
	return fmt.Errorf("item with key %s does not exist", key)
}
