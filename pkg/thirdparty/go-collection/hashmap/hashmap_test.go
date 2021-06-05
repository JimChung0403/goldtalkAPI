package hashmap

import (
	"testing"
	"goldtalkAPI/pkg/thirdparty/go-collection"
)

func TestHashMapNormal(t *testing.T) {

	var m collection.Map
	m = collection.Map(New())
	m.Put("123", 456)
	m.Put("456", "789")
	m.Put("abc", New())

	var v interface{}
	var found bool
	if v, found = m.Get("123"); !found {
		t.Log("123's value not found.")
		t.FailNow()
	}

	if v.(int) != 456 {
		t.Logf("Invalid 123's value: %d", v.(int))
		t.FailNow()
	}

	if v, found = m.Get("abc"); !found {
		t.Log("abc's value not found.")
		t.FailNow()
	}
}
