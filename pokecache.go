package main

import (
	"sync"
	"time"
)

// Esta sera la estructura del cache que usaremos para almacenar info en cache en vez de buscar
// cada vez que el usuario ponga un comando.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

// Dicho cache tendra la funcion add para anadir nuevos elementos
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

// Get para ver el valor de un elemento, con esto transformaremos ese []byte
// en informacion de la cual extraeremos lo que nos interese.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true

}

// Este metodo busca borrar el cache de informacion que no se haya buscado hace tiempo

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
