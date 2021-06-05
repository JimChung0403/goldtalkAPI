package hashmap

import (
	"sync"
)

const BUCKETS = 32

// MapItemBucket string:any
type MapItemBucket struct {
	items map[string]interface{}
	sync.RWMutex
}

// ConcurrentMap is an array of buckets
type ConcurrentMap []*MapItemBucket

// NewConcurrentMap returns a concurrently hashmap
func NewConcurrentMap() ConcurrentMap {
	m := make(ConcurrentMap, BUCKETS)
	for i := 0; i < BUCKETS; i++ {
		m[i] = &MapItemBucket{
			items: make(map[string]interface{}),
		}
	}
	return m
}

func (m ConcurrentMap) Clear() {

	for _, item := range m {
		item.Lock()
		item.items = make(map[string]interface{})
		item.Unlock()
	}
}

func (m ConcurrentMap) bucket(key string) *MapItemBucket {
	return m[hash(key)%BUCKETS]
}

// Set inserts key-value
func (m ConcurrentMap) Set(key string, value interface{}) {

	bucket := m.bucket(key)
	bucket.Lock()
	defer bucket.Unlock()
	bucket.items[key] = value

}

// SetIfAbsent returns whether first value of the key
func (m ConcurrentMap) SetIfAbsent(key string, value interface{}) bool {

	bucket := m.bucket(key)
	bucket.Lock()
	defer bucket.Unlock()
	_, ok := bucket.items[key]
	if !ok {
		bucket.items[key] = value
	}
	return !ok
}

func (m ConcurrentMap) Contains(key string) bool {

	bucket := m.bucket(key)
	bucket.RLock()
	defer bucket.RUnlock()
	_, ok := bucket.items[key]

	return ok
}

// Get returns value & success according to the given key
func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	bucket := m.bucket(key)
	bucket.RLock()
	defer bucket.RUnlock()
	value, ok := bucket.items[key]
	if value == nil {
		ok = false
	}
	return value, ok
}

func (m ConcurrentMap) Remove(key string) {

	bucket := m.bucket(key)
	bucket.Lock()
	defer bucket.Unlock()
	delete(bucket.items, key)
}

func (m ConcurrentMap) All() []interface{} {

	all := []interface{}{}
	for _, b := range m {
		b.RLock()
		for _, item := range b.items {
			all = append(all, item)
		}
		b.RUnlock()
	}
	return all
}

func (m ConcurrentMap) Keys() []string {
	keys := []string{}
	for _, b := range m {
		b.RLock()
		for key, _ := range b.items {
			keys = append(keys, key)
		}
		b.RUnlock()
	}
	return keys
}

func (m ConcurrentMap) Items() map[string]interface{} {
	ret := make(map[string]interface{})
	for _, b := range m {
		b.RLock()
		for key, value := range b.items {
			ret[key] = value
		}
		b.RUnlock()
	}
	return ret
}

func (m ConcurrentMap) Count() int {
	return len(m.Keys())
}

func hash(key string) uint32 {
	hashV := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hashV *= prime32
		hashV ^= uint32(key[i])
	}
	return hashV
}
