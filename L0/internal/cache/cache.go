package cache

import "sync"


type Cache struct {
    mu    sync.RWMutex
    cache map[string]interface{}
}


func NewCache() *Cache {
    return &Cache{
        cache: make(map[string]interface{}),
    }
}


func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    value, found := c.cache[key]
    return value, found
}


func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = value
}


func (c *Cache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.cache, key)
}