package services_utils

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

type LRUExpirableCache[K comparable, V any] struct {
	Size   int
	Expire time.Duration
	cache  *expirable.LRU[K, V]
	once   sync.Once
}

func (lrue *LRUExpirableCache[K, V]) Init() {
	lrue.once.Do(func() {
		if lrue.Expire == 0 {
			lrue.Expire = time.Minute * 5
			log.Print("LRUExpirableCache.Init: Expire is 0, setting to 5 minutes")
		}

		if lrue.Size == 0 {
			lrue.Size = 100
			log.Print("LRUExpirableCache.Init: Size is 0, setting to 100")
		}
		lrue.cache = expirable.NewLRU[K, V](lrue.Size, nil, lrue.Expire)
		log.Printf("LRUExpirableCache.Init: Size: %+v, Expire: %+v", lrue.Size, lrue.Expire)
	})
}

func (c *LRUExpirableCache[K, V]) Get(_ context.Context, key K) (V, bool) {
	v, ok := c.cache.Get(key)
	if ok {
		log.Printf("LRUCache.Get: %+v", key)
		return v, ok
	} else {
		log.Printf("LRUCache.Get ok false: %+v", key)
	}
	return v, ok
}

func (c *LRUExpirableCache[K, V]) Set(_ context.Context, key K, value V) {
	c.cache.Add(key, value)
	log.Printf("LRUCache.Set: (value :%+v, key: %+v)", value, key)
}

func (c *LRUExpirableCache[K, V]) Delete(_ context.Context, key K) bool {

	if c.cache.Contains(key) {
		c.cache.Remove(key)
		log.Printf("LRUCache.Delete: %+v", key)
		return true
	}
	return false
}

func (c *LRUExpirableCache[K, V]) Clear() {
	c.cache.Purge()
	log.Printf("LRUCache.Clear")
}
