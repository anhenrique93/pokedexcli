package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
