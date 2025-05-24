package internal

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mutex sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	ent := make(map[string]cacheEntry)
	cache := &Cache{
		entry: ent,
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Println("Add cache for: " + key)
	c.mutex.Lock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) (val []byte, wasFound bool) {
	fmt.Println("get cache for: " + key)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	ent, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return ent.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		go c.deleteEntryIfExpired(interval)
	}

}

func (c *Cache) deleteEntryIfExpired(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.entry) == 0 {
		return
	}

	for k, v := range c.entry {
		if time.Now().After(v.createdAt.Add(interval)) {
			delete(c.entry, k)
		}
	}
}
