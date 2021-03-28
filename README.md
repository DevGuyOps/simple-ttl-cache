# Simple TTL Cache
The Simple TTL Cache is an in memory cache with a set expiration time (TTL).

# Usage
```go
package main

import (
	"time"

	simplettlcache "github.com/GuySWatson/simple-ttl-cache"
)

func main() {
	// Setup
	simpleTTLCache := simplettlcache.SimpleTTLCache{}
	simpleTTLCache.Init(time.Second * 3)

	// Put object into cache
	simpleTTLCache.Put("color", "red")

	// Get object out of cache
	simpleTTLCache.Get("color")

	// Update object in the cache
	simpleTTLCache.Update("color", 1234)
}
```