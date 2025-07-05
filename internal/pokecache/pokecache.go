package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	entries map[string]cacheEntry
	ttl     time.Duration
	mu      sync.Mutex
}

func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		ttl:     ttl,
	}
	go c.reapLoop()
	return c
}

func NewClient(timeout time.Duration, cacheTTL time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecacheNewCache(cacheTTL),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if !exists {
		return nil, fmt.Errorf("cache entry for %s not found", key)
	}
	if time.Since(entry.createdAt) > c.ttl {
		delete(c.entries, key)
		return nil, fmt.Errorf("cache entry for %s has expired", key)
	}
	return entry.val, nil
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()
	for range ticker.C {
		now := time.Now()
		c.mu.Lock()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > c.ttl {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
