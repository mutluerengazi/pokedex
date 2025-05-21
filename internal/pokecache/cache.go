package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	items    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		items:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cp := make([]byte, len(val))
	copy(cp, val)

	c.items[key] = cacheEntry{
		createdAt: time.Now(),
		val:       cp,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.items[key]
	if !ok {
		return nil, false
	}

	cp := make([]byte, len(entry.val))
	copy(cp, entry.val)
	return cp, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, e := range c.items {
			if now.Sub(e.createdAt) > c.interval {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}
