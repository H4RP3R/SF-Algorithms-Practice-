// Используя мапу, реализуйте тип InMemoryCache, который позволит хранить значения в течение
// какого-то определённого времени (InMemoryCache должен реализовывать Cache interface):

package main

import "time"

var quit = make(chan bool)

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

func NewCacheEntry(val interface{}) CacheEntry {
	return CacheEntry{
		settledAt: time.Now(),
		value:     val,
	}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	storage map[string]CacheEntry
}

func (c *InMemoryCache) Set(key string, val interface{}) {
	entry := NewCacheEntry(val)
	c.storage[key] = entry
}

func (c *InMemoryCache) Get(key string) interface{} {
	return c.storage[key].value
}

func (c *InMemoryCache) storageTimer(delay time.Duration, storeTime time.Duration) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Second * delay)
			for key, entry := range c.storage {
				since := time.Since(entry.settledAt)
				if since*time.Second >= storeTime {
					delete(c.storage, key)
				}
			}
		}
	}
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	cache := InMemoryCache{
		storage: make(map[string]CacheEntry),
	}
	go cache.storageTimer(5, 10)
	return &cache
}
