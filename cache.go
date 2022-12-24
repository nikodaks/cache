package cache

type Cache struct {
	values map[string]interface{}
}

func (c *Cache) Set(key string, value interface{}) {
	c.values[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.values[key]
}

func (c *Cache) Delete(key string) {
	delete(c.values, key)
}

func New() *Cache {
	return &Cache{
		make(map[string]interface{}),
	}
}
