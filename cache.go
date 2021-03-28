package simplettlcache

import (
	"sync"
	"time"
)

type SimpleTTLCache struct {
	assets map[string]*simpleTTLItem
	mutex  sync.Mutex
}

type simpleTTLItem struct {
	value      interface{}
	lastAccess time.Time
}

func (s *SimpleTTLCache) Init(maxTTL time.Duration) {
	s.assets = make(map[string]*simpleTTLItem)

	go func() {
		for now := range time.Tick(time.Second) {
			s.mutex.Lock()

			for k, v := range s.assets {
				if now.Sub(v.lastAccess) > maxTTL {
					delete(s.assets, k)
				}
			}

			s.mutex.Unlock()
		}
	}()
}

func (s *SimpleTTLCache) Len() int {
	return len(s.assets)
}

func (s *SimpleTTLCache) Put(k string, value interface{}) {
	s.mutex.Lock()

	item, ok := s.assets[k]
	if !ok {
		item = &simpleTTLItem{
			value: value,
		}
		s.assets[k] = item
	}

	item.lastAccess = time.Now()

	s.mutex.Unlock()
}

func (s *SimpleTTLCache) Get(k string) (v interface{}) {
	s.mutex.Lock()

	if item, ok := s.assets[k]; ok {
		v = item.value
		item.lastAccess = time.Now()
	}

	s.mutex.Unlock()
	return
}

func (s *SimpleTTLCache) Update(k string, value interface{}) {
	s.mutex.Lock()

	item, ok := s.assets[k]
	if ok {
		item.value = value
		item.lastAccess = time.Now()
	}

	s.mutex.Unlock()
}
