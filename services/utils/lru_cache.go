package services_utils

import (
	"context"
	"errors"
	"log"
	"sync"

	dataloader "github.com/graph-gophers/dataloader/v7"

	lru "github.com/hashicorp/golang-lru"
)

var once sync.Once
var ctx *lru.ARCCache

// Cache implements the dataloader.Cache interface
type LRUCache[K comparable, V any] struct {
	Size int
}

func (lruc *LRUCache[K, V]) ARCCache() (c *lru.ARCCache, err error) {
	once.Do(func() {
		ctx, err = lru.NewARC(lruc.Size)
		if err != nil {
			return
		}
	})

	if ctx == nil {
		return nil, errors.New("LRUCache.GetARC: ctx is nil")
	}

	return ctx, nil
}

// Get gets an item from the cache
func (c *LRUCache[K, V]) Get(_ context.Context, key K) (dataloader.Thunk[V], bool) {
	arcCache, _ := c.ARCCache()

	v, ok := arcCache.Get(key)
	if ok {
		log.Printf("LRUCache.Get: %+v", key)
		return v.(dataloader.Thunk[V]), ok
	} else {
		log.Printf("LRUCache.Get ok false: %+v", key)
	}
	return nil, ok
}

// Set sets an item in the cache
func (c *LRUCache[K, V]) Set(_ context.Context, key K, value dataloader.Thunk[V]) {
	arcCache, _ := c.ARCCache()
	arcCache.Add(key, value)
	log.Printf("LRUCache.Set: (value :%+v, key: %+v)", value, key)
}

// Delete deletes an item in the cache
func (c *LRUCache[K, V]) Delete(_ context.Context, key K) bool {
	arcCache, _ := c.ARCCache()

	if arcCache.Contains(key) {
		arcCache.Remove(key)
		log.Printf("LRUCache.Delete: %+v", key)
		return true
	}
	return false
}

// Clear clears the cache
func (c *LRUCache[K, V]) Clear() {
	arcCache, _ := c.ARCCache()
	arcCache.Purge()
	log.Printf("LRUCache.Clear")
}
