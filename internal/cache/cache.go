package cache

import (
	"runtime"
	"sync"
	"time"
)

type cacheValue struct {
	Value      interface{}
	Expiration int64
}

type Cache struct {
	cache      map[string]cacheValue
	sync       sync.RWMutex
	expiration time.Duration
	cleaner    *cleaner
}

type CacheSettings struct {
	Expiration      time.Duration
	CleanupInterval time.Duration
}

func NewCache(settings CacheSettings) *Cache {
	c := &Cache{
		cache:      make(map[string]cacheValue),
		sync:       sync.RWMutex{},
		expiration: settings.Expiration,
	}

	if settings.CleanupInterval > 0 {
		runCleaner(c, settings.CleanupInterval)
		runtime.SetFinalizer(c, stopCleaner)
	}

	return c
}

func (c *Cache) Get(key string) (cacheValue, bool) {
	c.sync.RLock()
	defer c.sync.RUnlock()

	val, ok := c.cache[key]
	if !ok {
		return cacheValue{}, false
	}
	if val.Expiration < time.Now().UnixNano() {
		delete(c.cache, key)
		return cacheValue{}, false
	}
	return val, ok
}

func (c *Cache) Set(key string, val interface{}) {
	c.sync.Lock()
	defer c.sync.Unlock()
	now := time.Now()
	c.cache[key] = cacheValue{
		Value:      val,
		Expiration: now.Add(c.expiration).UnixNano(),
	}
}

func (c *Cache) Delete(key string) {
	c.sync.Lock()
	defer c.sync.Unlock()
	delete(c.cache, key)
}

func (c *Cache) flushExpired() {
	c.sync.Lock()
	defer c.sync.Unlock()
	for k, v := range c.cache {
		if v.Expiration < time.Now().UnixNano() {
			delete(c.cache, k)
		}
	}
}

type cleaner struct {
	interval time.Duration
	stop     chan bool
}

func (cln *cleaner) run(c *Cache) {
	ticker := time.NewTicker(cln.interval)
	for {
		select {
		case <-cln.stop:
			ticker.Stop()
			return
		}
	}
}

func runCleaner(c *Cache, interval time.Duration) {
	cleaner := &cleaner{
		interval: interval,
		stop:     make(chan bool),
	}

	c.cleaner = cleaner

	go cleaner.run(c)
}

func stopCleaner(c *Cache) {
	c.cleaner.stop <- true
}
