package redis

import (
	"testing"
	//"container/list"
	"fmt"
	"goldtalkAPI/pkg/thirdparty/go-collection/hashmap"
	"time"
)

type GOBEntry struct {
	Id   int32
	Name string
	M    map[string]interface{}
	a    int
}

func TestSet(t *testing.T) {

	InitRedisPool("172.16.235.48:6380", 1024, "test")

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

func TestINCR(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "incr"
	t.Log(INCR(key))
	t.Log(INCR(key))
	t.Log(INCR(key))
}

func TestSetExpireAt(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "incr"
	SetExpireAt(key, time.Now().Unix()+100)
}

func TestSetExpire(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "incr"
	SetExpire(key, 100)
}

func TestSetStr(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "setStr"
	t.Log(SetStr(key, key, 100))

	t.Log(SetStr(key+"1", key, 0))
}

func TestGetStr(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "setStr"
	t.Log(GetStr(key))
}

func TestHDel(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "hdel"
	sns := map[string]string{
		"q": "1",
		"w": "1",
		"e": "1",
	}
	t.Error(Hmset2(key, sns))
	t.Log(HGetAllMap(key))
	t.Log(HDel(key, "w"))
	t.Log(HGetAllMap(key))
}

func TestHlen(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "hlen"
	sns := map[string]string{
		"q": "1",
		"w": "1",
		"e": "1",
	}
	t.Error(Hmset2(key, sns))
	t.Log(HGetAllMap(key))
	t.Log(HLen(key))
	t.Log(HDel(key, "w"))
	t.Log(HGetAllMap(key))
	t.Log(HLen(key))
}

func TestSetPersist(t *testing.T) {
	InitRedisPool("172.16.235.144:6380", 1024, "")
	key := "persist"
	sns := map[string]string{
		"q": "1",
		"w": "1",
		"e": "1",
	}
	t.Log(Hmset2(key, sns))
	time.Sleep(time.Second*5)
	SetPersist(key)
}
