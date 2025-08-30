package chanjs

import (
    "sync"
    "time"
)

type DedupCache struct {
    mu       sync.Mutex
    cache    map[string]time.Time
    maxSize  int
    ttl      time.Duration
}

func NewDedupCache(maxSize int, ttl time.Duration) *DedupCache {
    return &DedupCache{
        cache:   make(map[string]time.Time),
        maxSize: maxSize,
        ttl:     ttl,
    }
}

func (d *DedupCache) IsDuplicate(msgID string) bool {
    d.mu.Lock()
    defer d.mu.Unlock()
    now := time.Now()
    if ts, exists := d.cache[msgID]; exists {
        if now.Sub(ts) < d.ttl {
            return true
        }
    }
    d.cache[msgID] = now
    if len(d.cache) > d.maxSize {
        // lógica simple: eliminar el más antiguo
        oldestID, oldestTime := "", now
        for id, t := range d.cache {
            if t.Before(oldestTime) {
                oldestTime, oldestID = t, id
            }
        }
        delete(d.cache, oldestID)
    }
    return false
}

