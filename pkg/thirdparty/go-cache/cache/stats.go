package cache

import (
	"sync"
	"sync/atomic"
)

type statsAccessor interface {
	HitCount() uint64
	MissCount() uint64
	LookupCount() uint64
	HitRate() float64
}

type stats struct {
	hitCount  uint64
	missCount uint64
}

// IncrementHitCount increment hit count
func (st *stats) IncrementHitCount() uint64 {
	return atomic.AddUint64(&st.hitCount, 1)
}

// IncrementMissCount increment miss count
func (st *stats) IncrementMissCount() uint64 {
	return atomic.AddUint64(&st.missCount, 1)
}

// HitCount returns hit count
func (st *stats) HitCount() uint64 {
	return atomic.LoadUint64(&st.hitCount)
}

// MissCount returns miss count
func (st *stats) MissCount() uint64 {
	return atomic.LoadUint64(&st.missCount)
}

// LookupCount returns lookup count
func (st *stats) LookupCount() uint64 {
	return st.HitCount() + st.MissCount()
}

// HitRate returns rate for cache hitting
func (st *stats) HitRate() float64 {
	hc, mc := st.HitCount(), st.MissCount()
	total := hc + mc
	if total == 0 {
		return 0.0
	}
	return float64(hc) / float64(total)
}

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group represents a class of work
type Group struct {
	cache Cache
	mu    sync.Mutex
	m     map[interface{}]*call
}

// Do executes and returns the results of the given function
func (g *Group) Do(key interface{}, fn func() (interface{}, error), isWait bool) (interface{}, bool, error) {
	g.mu.Lock()
	v, err := g.cache.get(key, true)
	if err == nil {
		g.mu.Unlock()
		return v, false, nil
	}
	if g.m == nil {
		g.m = make(map[interface{}]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		if !isWait {
			return nil, false, KeyNotFoundError
		}
		c.wg.Wait()
		return c.val, false, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()
	if !isWait {
		go g.call(c, key, fn)
		return nil, false, KeyNotFoundError
	}
	v, err = g.call(c, key, fn)
	return v, true, err
}

func (g *Group) call(c *call, key interface{}, fn func() (interface{}, error)) (interface{}, error) {
	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
