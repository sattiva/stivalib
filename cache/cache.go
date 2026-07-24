package cache

import (
	"sync"
	"time"
)

type item struct {
	val interface{}
	exp time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]item
}

func New(ttl time.Duration) *Cache {
	c := &Cache{items: make(map[string]item)}
	go c.janitor(ttl)
	return c
}

func (c *Cache) Set(key string, val interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = item{val: val, exp: time.Now().Add(ttl)}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	it, ok := c.items[key]
	if !ok || time.Now().After(it.exp) {
		return nil, false
	}
	return it.val, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *Cache) janitor(interval time.Duration) {
	for range time.Tick(interval) {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.items {
			if now.After(v.exp) {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}
