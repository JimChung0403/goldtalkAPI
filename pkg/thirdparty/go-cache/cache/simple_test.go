package cache

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

const (
	simple_loops = 10
)

func forkSimpleCacheForTest(size int) Cache {
	return New(size).Simple().EvictedFunc(func(key, value interface{}) {
		fmt.Printf("key:%v value:%v evicted.\n", key, value)
	}).Expiration(1 * time.Second).Build()
}

func TestSimpleGetMany(t *testing.T) {

	wg := sync.WaitGroup{}
	c := forkSimpleCacheForTest(simple_loops)
	for i := 0; i < simple_loops; i++ {
		wg.Add(1)
		c.Set(strconv.Itoa(i), i)
	}

	go func() {

		time.Sleep(200 * time.Millisecond)
		for i := 0; i < simple_loops; i++ {

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
	//t.Log(fmt.Sprintf("hit:%v miss:%v lookup:%v rate:%v", c.HitCount(), c.MissCount(), c.LookupCount(), c.HitRate()))
}

func TestSimpleGetExpire(t *testing.T) {

	c := forkSimpleCacheForTest(10)
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

	if c.HitCount() != 1 {
		t.Error("hitcount not correct")
		t.FailNow()
	}

	if c.MissCount() != 1 {
		t.Error("misscount not correct")
		t.FailNow()
	}

	if c.LookupCount() != 2 {
		t.Error("lookupcount not correct")
		t.FailNow()
	}
}
