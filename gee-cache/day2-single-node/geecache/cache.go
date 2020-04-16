package geecache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mutex      sync.Mutex
	lruCache   *lru.Cache
	cacheBytes int64
}

func (cache *cache) add(key string, value ByteView) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if cache.lruCache == nil {
		cache.lruCache = lru.New(cache.cacheBytes, nil)
	}
	cache.lruCache.Add(key, value)
}

func (cache *cache) get(key string) (value ByteView, success bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if cache.lruCache == nil {
		return
	}

	if val, success := cache.lruCache.Get(key); success {
		return val.(ByteView), success
	}

	return
}
