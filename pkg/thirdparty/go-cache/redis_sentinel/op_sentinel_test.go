package redis

import (
	"fmt"
	"goldtalkAPI/pkg/thirdparty/go-collection/hashmap"
	"testing"
	"time"
)

type GOBEntry struct {
	Id   int32
	Name string
	M    map[string]interface{}
	a    int
}

func TestSet(t *testing.T) {

	addrs := []string{"172.16.235.145:26381", "172.16.235.145:26382", "172.16.235.145:26383"}
	InitSentinelRedisPool(addrs, "mymaster", 20, "test")

	bb := &GOBEntry{
		Id:   123,
		Name: "dellinger",
		M:    make(map[string]interface{}),
	}

	m := hashmap.NewConcurrentMap()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%04d", i)
		//bb.M[fmt.Sprintf("key%04d", i)] = i

		//bb.MM.PushBack(fmt.Sprintf("key%04d", i))
		m.Set(key, i)
	}

	bb.M = m.Items()

	if err := Set("ddd", bb, 30*time.Second); err != nil {
		t.Error(err)
		t.Fail()
	}

	b := new(GOBEntry)
	if err := Get("ddd", b); err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Logf("%#v", b)

	if err := Delete("ddd"); err != nil {
		t.Error(err)
		t.Fail()
	}
}
