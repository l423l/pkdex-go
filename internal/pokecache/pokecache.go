package pokecache

import (
 "time"
 "sync"
)

type cacheEntry struct {
  createdAt time.Time
  val []byte
}

type Cache struct {
  caches map[string]cacheEntry
  mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
  ch := &Cache{
    caches: make(map[string]cacheEntry),
  }
  go ch.reapLoop(interval)
  return ch
}

func (ch *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
  for range ticker.C {
    ch.mu.Lock()
    for key, cache := range ch.caches {
      if interval < time.Since(cache.createdAt) {
        delete(ch.caches, key)
      }
    }
  ch.mu.Unlock()
  }
}

func (ch *Cache) Add(key string, val []byte) {
  ch.mu.Lock()
  defer ch.mu.Unlock()

  ch.caches[key] = cacheEntry{
    createdAt: time.Now(),
    val: val,
  }
}

func (ch *Cache) Get(key string) ([]byte, bool) {
  ch.mu.Lock()
  defer ch.mu.Unlock()

  cache, ok := ch.caches[key]
  if !ok {
    return nil, false
  }
  return cache.val, true
}
