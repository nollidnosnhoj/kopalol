package cache

import (
	"bytes"
	"sync"
	"time"
)

var EXPIRATION = 24 * time.Hour

type cacheValue struct {
	val         bytes.Buffer
	contentType string
	expires     int64
}

type Cache struct {
	cache map[string]cacheValue
	sync  sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]cacheValue),
		sync:  sync.RWMutex{},
	}
}

func (c *Cache) Has(key string) bool {
	c.sync.RLock()
	defer c.sync.RUnlock()
	val, ok := c.cache[key]
	if !ok {
		return false
	}
	if val.expires < time.Now().Unix() {
		delete(c.cache, key)
		return false
	}
	return true
}

func (c *Cache) Get(key string) (bytes.Buffer, string, bool) {
	c.sync.RLock()
	defer c.sync.RUnlock()
	val, ok := c.cache[key]
	if !ok {
		return bytes.Buffer{}, "", false
	}
	if val.expires < time.Now().Unix() {
		delete(c.cache, key)
		return bytes.Buffer{}, "", false
	}
	return val.val, val.contentType, ok
}

func (c *Cache) Set(key string, val bytes.Buffer, contentType string) {
	c.sync.Lock()
	defer c.sync.Unlock()
	now := time.Now()
	c.cache[key] = cacheValue{
		val:         val,
		contentType: contentType,
		expires:     now.Add(EXPIRATION).Unix(),
	}
}

func (c *Cache) Delete(key string) {
	c.sync.Lock()
	defer c.sync.Unlock()
	delete(c.cache, key)
}
