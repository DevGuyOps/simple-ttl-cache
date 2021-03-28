package simplettlcache

import (
	"testing"
	"time"
)

func TestSimpleTTLCache(t *testing.T) {
	key := "a1"

	// Setup cache for test
	simpleTTLCache := SimpleTTLCache{}
	simpleTTLCache.Init(time.Second * 3)

	// Put item in cache and check it exists
	simpleTTLCache.Put(key, "red")
	if simpleTTLCache.Get(key) != "red" {
		t.Error("Item missing from cache")
	}

	// Sleep to test TTL works
	time.Sleep(time.Second * 4)
	if simpleTTLCache.Get(key) != nil {
		t.Error("Should be empty")
	}
}

func TestSimpleTTLCacheLen(t *testing.T) {
	key := "b2"

	// Setup cache for test
	simpleTTLCache := SimpleTTLCache{}
	simpleTTLCache.Init(time.Second * 3)

	// Put item in cache and check length
	simpleTTLCache.Put(key, "apple")
	if simpleTTLCache.Len() != 1 {
		t.Error("Len of cache is wrong")
	}
}

func TestSimpleTTLCacheUpdate(t *testing.T) {
	key := "b2"

	// Setup cache for test
	simpleTTLCache := SimpleTTLCache{}
	simpleTTLCache.Init(time.Second * 3)

	// Put item in cache and check it exists
	simpleTTLCache.Put(key, "apple")
	if simpleTTLCache.Get(key) != "apple" {
		t.Error("Item missing from cache")
	}

	// Update item in cache and check it changed
	simpleTTLCache.Update(key, "banana")
	if simpleTTLCache.Get(key) != "banana" {
		t.Error("Item update failed")
	}
}
