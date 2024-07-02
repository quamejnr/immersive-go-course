package cache

import (
	"sync"
	"time"
)

type cacheEntryStats struct {
	reads        int
	writes       int
	lastAccessed time.Time
}

type cacheEntry[V any] struct {
	value V
	stats cacheEntryStats
}

type cache[T comparable, V any] struct {
	sync.Mutex
	limit           int
	totalReads      int
	successfulReads int
	totalWrites     int
	entries         map[T]*cacheEntry[V]
}

func NewCache[T comparable, V any](entryLimit int) *cache[T, V] {
	return &cache[T, V]{
		limit:   entryLimit,
		entries: make(map[T]*cacheEntry[V], entryLimit),
	}
}

// Returns the key of the entry that is least accessed
func (c *cache[T, V]) LRU() T {
	lru := struct {
		minTime time.Time
		key     T
	}{}
	isFirst := true
	for k, v := range c.entries {
		if isFirst || v.stats.lastAccessed.Before(lru.minTime) {
			lru.minTime = v.stats.lastAccessed
			lru.key = k
			isFirst = false
		}
	}
	return lru.key
}

// Put entries in cache and returns true if entry already existed
// If entry already present, replace entry.
func (c *cache[T, V]) Put(key T, val V) bool {
	if c.limit == 0 {
		return false
	}

	c.Lock()
	defer c.Unlock()

	e, present := c.entries[key]

	if !present {
		// remove least recently used if entries is filled to the limit
		if len(c.entries) == c.limit {
			k := c.LRU()
			delete(c.entries, k)
		}
		e = &cacheEntry[V]{}
		c.entries[key] = e
	}
	e.value = val
	e.stats.lastAccessed = time.Now()
	e.stats.writes++
	c.totalWrites++
	return present
}

// Get returns the value of the key passed and a boolean that shows whether the entry existed
// If the entry did not exist, the boolean is false and value is nil.
func (c *cache[T, V]) Get(key T) (*V, bool) {
	c.Lock()
	defer c.Unlock()
	e, ok := c.entries[key]
	c.totalReads++
	if !ok {
		return nil, false
	}
	c.successfulReads++
	e.stats.reads++
	e.stats.lastAccessed = time.Now()
	return &e.value, ok
}

// func NewCache[K comparable, V any](entryLimit int) Cache[K, V] { ... }
// Put adds the value to the cache, and returns a boolean to indicate whether a value already existed in the cache for that key.
// If there was previously a value, it replaces that value with this one.
// Any Put counts as a refresh in terms of LRU tracking.
// func (c *Cache[K, V]) Put(key K, value V) bool { ... }

// Get returns the value assocated with the passed key, and a boolean to indicate whether a value was known or not. If not, nil is returned as the value.
