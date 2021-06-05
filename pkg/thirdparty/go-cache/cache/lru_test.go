package cache

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

const (
	lru_loops = 10
)

func forkLRUCacheForTest(size int) Cache {
	return New(size).LRU().EvictedFunc(func(key, value interface{}) {
		fmt.Printf("key:%v value:%v evicted.\n", key, value)
	}).Expiration(1 * time.Second).Build()
}

func TestLRUGetMany(t *testing.T) {

	wg := sync.WaitGroup{}
	c := forkLRUCacheForTest(lru_loops)
	for i := 0; i < lru_loops; i++ {
		wg.Add(1)
		c.Set(strconv.Itoa(i), i)
	}

	go func() {

		time.Sleep(200 * time.Millisecond)
		for i := 0; i < lru_loops; i++ {

			defer wg.Add(-1)
			obj, err := c.Get(strconv.Itoa(i))
			if err != nil {
				t.Errorf("fail to get value:%v", i)
			} else if i != obj.(int) {
				t.Errorf("index: %v not correct", i)
				t.FailNow()
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
}

func TestLRUGetExpire(t *testing.T) {

	c := forkLRUCacheForTest(10)
	c.SetWithExpire("k1", "v1", time.Second)
	time.Sleep(500 * time.Millisecond)

	if v, err := c.Get("k1"); err != nil {
		t.Error("fail to get k1.", err)
		t.FailNow()
	} else if "v1" != v.(string) {

		t.Error("value mismatch.", err)
		t.FailNow()
	}

	time.Sleep(500 * time.Millisecond)
	if _, err := c.Get("k1"); err == nil {
		t.Error("k1 should be expired but not.")
		t.FailNow()
	}
}

func TestLRUNormal(t *testing.T) {

	c := forkLRUCacheForTest(3)
	c.Set("k1", "v1")
	c.Set("k2", "v2")
	c.Set("k3", "v3")

	c.Get("k1")
	c.Get("k1")
	c.Get("k1")
	c.Get("k2")
	c.Get("k2")
	c.Get("k3")

	c.Set("k4", "v4")
	if _, err := c.Get("k1"); err == nil {
		t.Error("k1 should be evicted")
		t.FailNow()
	}

	t.Log(fmt.Sprintf("hit:%v miss:%v lookup:%v rate:%v", c.HitCount(), c.MissCount(), c.LookupCount(), c.HitRate()))
}
