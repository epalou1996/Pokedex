package main

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true

}

func (c *Cache) reapLoop(duration time.Duration) {
	keysToDelete := make([]string, len(c.cache))
	currentTime := time.Now()

	c.mu.RLock()
	for key, entry := range c.cache {
		timeElapsed := currentTime.Sub(entry.createdAt)
		if timeElapsed < duration {
			keysToDelete = append(keysToDelete, key)
		}
	}
	c.mu.RUnlock()
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, key := range keysToDelete {
		delete(c.cache, key)
	}
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}

	go cache.reapLoop(duration)
	return cache
}
