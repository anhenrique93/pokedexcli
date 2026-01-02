package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.data[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.data[key]
	return val.val, ok
}

func (c *Cache) cleanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for k, entry := range c.data {
		if now.Sub(entry.createdAt) > c.interval {
			delete(c.data, k)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.cleanUp()
	}
}
