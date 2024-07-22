package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      string
	Expiration int64
}

type LRUCache struct {
	capacity int
	items    map[string]*CacheItem
	mutex    sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*CacheItem),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found || item.Expiration <= time.Now().Unix() {
		return "", false
	}
	return item.Value, true
}

func (c *LRUCache) Set(key string, value string, expiration int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items[key] = &CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *LRUCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}
