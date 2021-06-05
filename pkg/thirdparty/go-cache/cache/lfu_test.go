package cache

import (
	"fmt"
	logger "scm.tutorabc.com/Research/GoBank/goutil/log"
	"strconv"
	"sync"
	"testing"
	"time"
)

const (
	lfu_loops = 10
)

func forkLFUCacheForTest(size int) Cache {
	return New(size).LFU().EvictedFunc(func(key, value interface{}) {
		fmt.Printf("key:%v value:%v evicted.\n", key, value)
	}).Expiration(1 * time.Second).Build()
}

func TestLFUGetMany(t *testing.T) {

	wg := sync.WaitGroup{}
	c := forkLFUCacheForTest(lfu_loops)
	for i := 0; i < lfu_loops; i++ {
		wg.Add(1)
		c.Set(strconv.Itoa(i), i)
	}

	go func() {

		time.Sleep(200 * time.Millisecond)
		for i := 0; i < lfu_loops; i++ {

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

func TestLFUGetExpire(t *testing.T) {

	c := forkLFUCacheForTest(10)
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

func TestLFUNormal(t *testing.T) {

	c := forkLFUCacheForTest(3)
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
	if _, err := c.Get("k3"); err == nil {
		t.Error("k3 should be evicted")
		t.FailNow()
	}
}

func TestCache(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(3)
	cache := New(1024).LRU().Expiration(10 * time.Minute).Build()
	go func() {
		for i := 0; i < 100; i++ {
			cache.Set(fmt.Sprintf("key%d", i), "123")
			time.Sleep(99 * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			cache.Remove(fmt.Sprintf("key%d", i))
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(100 * time.Millisecond)
			if v, err := cache.Get(fmt.Sprintf("key%d", i)); err != nil {
				logger.Errorf("fail to get key%d", i)
			} else {
				logger.Infof("got key%d:%v", i, v)
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
