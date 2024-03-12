package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap   map[string]cacheEntry
	cacheMutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheMap:   make(map[string]cacheEntry),
		cacheMutex: &sync.Mutex{},
	}
	go c.startEviction(interval)
	return c
}

func (c *Cache) startEviction(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			c.evict()
		}
	}
}

func (c *Cache) evict() {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	for key, entry := range c.cacheMap {
		if time.Since(entry.createdAt) > 5*time.Minute {
			delete(c.cacheMap, key)
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	entry, exists := c.cacheMap[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}
